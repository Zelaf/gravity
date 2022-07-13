package dns

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"beryju.io/gravity/pkg/roles"
	"beryju.io/gravity/pkg/roles/dns/types"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Zone struct {
	Name string `json:"-"`

	Authoritative  bool                `json:"authoritative"`
	HandlerConfigs []map[string]string `json:"handlerConfigs"`
	DefaultTTL     uint32              `json:"defaultTTL"`

	h []Handler

	records         map[string]map[string]*Record
	recordsSync     sync.RWMutex
	recordsWatchCtx context.CancelFunc

	etcdKey string
	inst    roles.Instance
	log     *log.Entry
}

func (z *Zone) metrics(dur time.Duration, q *dns.Msg, h Handler, rep *dns.Msg) {
	for _, question := range q.Question {
		dnsQueries.WithLabelValues(
			dns.TypeToString[question.Qtype],
			h.Identifier(),
			z.Name,
		).Inc()
		dnsQueryDuration.WithLabelValues(
			dns.TypeToString[question.Qtype],
			h.Identifier(),
			z.Name,
		).Observe(float64(dur.Milliseconds()))
	}
}

func (z *Zone) resolve(w dns.ResponseWriter, r *dns.Msg) {
	for _, handler := range z.h {
		z.log.WithField("handler", handler.Identifier()).Trace("sending request to handler")
		start := time.Now()
		handlerReply := handler.Handle(NewFakeDNSWriter(w), r)
		finish := time.Since(start)
		if handlerReply != nil {
			z.log.WithField("handler", handler.Identifier()).Trace("returning reply from handler")
			handlerReply.SetReply(r)
			w.WriteMsg(handlerReply)
			z.metrics(finish, r, handler, handlerReply)
			return
		}
		z.log.WithField("handler", handler.Identifier()).Trace("no reply, trying next handler")
	}
	z.log.Debug("No handler has a reply, fallback back to NX")
	fallback := new(dns.Msg)
	fallback.SetReply(r)
	fallback.SetRcode(r, dns.RcodeNameError)
	w.WriteMsg(fallback)
}

func (r *DNSRole) zoneFromKV(raw *mvccpb.KeyValue) (*Zone, error) {
	z := Zone{
		DefaultTTL:  3600,
		inst:        r.i,
		h:           make([]Handler, 0),
		records:     make(map[string]map[string]*Record),
		recordsSync: sync.RWMutex{},
	}
	err := json.Unmarshal(raw.Value, &z)
	if err != nil {
		return nil, err
	}
	prefix := r.i.KV().Key(types.KeyRole, types.KeyZones, "")
	z.Name = strings.TrimPrefix(string(raw.Key), prefix)
	// Get full etcd key without leading slash since this usually gets passed to Instance Key()
	z.etcdKey = string(raw.Key)[1:]
	z.log = r.log.WithField("zone", z.Name)

	if len(z.HandlerConfigs) < 1 {
		r.log.Trace("No handler defined, defaulting to etcd")
		z.HandlerConfigs = append(z.HandlerConfigs, map[string]string{
			"type": "etcd",
		})
	}

	for _, handlerCfg := range z.HandlerConfigs {
		t := handlerCfg["type"]
		var handler Handler
		var err error
		switch t {
		case "forward_blocky":
			handler, err = NewBlockyForwarder(&z, handlerCfg)
		case "forward_ip":
			handler = NewIPForwarderHandler(&z, handlerCfg)
		case "etcd":
			handler = NewEtcdHandler(&z, handlerCfg)
		case "memory":
			handler = NewMemoryHandler(&z, handlerCfg)
		default:
			r.log.WithField("type", t).Warning("invalid forwarder type")
		}
		if err != nil {
			z.log.WithError(err).Warning("failed to setup handler")
			continue
		}
		z.h = append(z.h, handler)
	}

	// start watching all records in this zone, in case etcd goes down
	go z.watchZoneRecords()

	return &z, nil
}

func (z *Zone) watchZoneRecords() {
	evtHandler := func(ev *clientv3.Event) {
		z.recordsSync.Lock()
		defer z.recordsSync.Unlock()
		rec, err := z.recordFromKV(ev.Kv)
		if _, ok := z.records[rec.recordKey]; !ok {
			z.records[rec.recordKey] = make(map[string]*Record)
		}
		if ev.Type == clientv3.EventTypeDelete {
			delete(z.records[rec.recordKey], rec.uid)
			dnsRecordsMetric.WithLabelValues(z.Name).Dec()
		} else {
			// Check if the record parsed above actually was parsed correctly,
			// we don't care for that when removing, but prevent adding
			// empty records
			if err != nil {
				return
			}
			if _, ok := z.records[rec.recordKey][rec.uid]; !ok {
				dnsRecordsMetric.WithLabelValues(z.Name).Inc()
			}
			z.records[rec.recordKey][rec.uid] = rec
		}
	}
	ctx, canc := context.WithCancel(context.Background())
	z.recordsWatchCtx = canc

	prefix := "/" + z.etcdKey + "/"

	records, err := z.inst.KV().Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		z.log.WithError(err).Warning("failed to list initial records")
		time.Sleep(5 * time.Second)
		z.watchZoneRecords()
		return
	}
	for _, record := range records.Kvs {
		evtHandler(&clientv3.Event{
			Type: mvccpb.PUT,
			Kv:   record,
		})
	}

	watchChan := z.inst.KV().Watch(
		ctx,
		prefix,
		clientv3.WithPrefix(),
		clientv3.WithProgressNotify(),
	)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			go evtHandler(event)
		}
	}
}

func (z *Zone) StopWatchingRecords() {
	if z.recordsWatchCtx != nil {
		z.recordsWatchCtx()
	}
}

func (z *Zone) put() error {
	raw, err := json.Marshal(&z)
	if err != nil {
		return err
	}

	_, err = z.inst.KV().Put(
		context.TODO(),
		z.inst.KV().Key(
			types.KeyRole,
			types.KeyZones,
			z.Name,
		),
		string(raw),
	)
	if err != nil {
		return err
	}
	return nil
}

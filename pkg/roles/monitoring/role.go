package monitoring

import (
	"context"
	"net/http"
	_ "unsafe"

	"beryju.io/gravity/pkg/extconfig"
	"beryju.io/gravity/pkg/roles"
	"beryju.io/gravity/pkg/roles/api"
	apitypes "beryju.io/gravity/pkg/roles/api/types"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/swaggest/rest/web"
)

type Role struct {
	m      *mux.Router
	log    *log.Entry
	i      roles.Instance
	ctx    context.Context
	cfg    *RoleConfig
	server *http.Server
}

//go:linkname blockyReg github.com/0xERR0R/blocky/metrics.reg
var blockyReg = prometheus.NewRegistry()

func New(instance roles.Instance) *Role {
	mux := mux.NewRouter()
	r := &Role{
		log: instance.Log(),
		i:   instance,
		m:   mux,
	}
	r.m.Use(api.NewRecoverMiddleware(r.log))
	r.m.Use(api.NewLoggingMiddleware(r.log, nil))
	r.m.Path("/healthz/live").HandlerFunc(r.HandleHealthLive)
	r.m.Path("/metrics").HandlerFunc(r.HandleMetrics)
	r.i.AddEventListener(apitypes.EventTopicAPIMuxSetup, func(ev *roles.Event) {
		svc := ev.Payload.Data["svc"].(*web.Service)
		svc.Get("/api/v1/roles/monitoring", r.APIRoleConfigGet())
		svc.Post("/api/v1/roles/monitoring", r.APIRoleConfigPut())
	})
	return r
}

func (r *Role) HandleHealthLive(w http.ResponseWriter, re *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (r *Role) HandleMetrics(w http.ResponseWriter, re *http.Request) {
	promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		DisableCompression: true,
	})).ServeHTTP(w, re)
	promhttp.HandlerFor(blockyReg, promhttp.HandlerOpts{
		DisableCompression: true,
	}).ServeHTTP(w, re)
}

func (r *Role) Start(ctx context.Context, config []byte) error {
	r.ctx = ctx
	r.cfg = r.decodeRoleConfig(config)
	listen := extconfig.Get().Listen(r.cfg.Port)
	r.log.WithField("listen", listen).Info("starting monitoring Server")
	r.server = &http.Server{
		Addr:    listen,
		Handler: r.m,
	}
	go func() {
		err := r.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			r.log.WithError(err).Warning("failed to listen")
		}
	}()
	return nil
}

func (r *Role) Stop() {
	if r.server != nil {
		r.server.Shutdown(r.ctx)
	}
}

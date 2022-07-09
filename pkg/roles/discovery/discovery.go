package discovery

import (
	"beryju.io/ddet/pkg/extconfig"
	"beryju.io/ddet/pkg/roles"
	apitypes "beryju.io/ddet/pkg/roles/api/types"
	"beryju.io/ddet/pkg/roles/discovery/types"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type DiscoveryRole struct {
	log *log.Entry
	i   roles.Instance
	cfg *DiscoveryRoleConfig
}

func New(instance roles.Instance) *DiscoveryRole {
	r := &DiscoveryRole{
		log: instance.GetLogger().WithField("role", types.KeyRole),
		i:   instance,
	}
	r.i.AddEventListener(apitypes.EventTopicAPIMuxSetup, func(ev *roles.Event) {
		mux := ev.Payload.Data["mux"].(*chi.Mux)
		mux.Post("/v0/discovery/apply", r.apiHandlerApply)
	})
	return r
}

func (r *DiscoveryRole) Start(config []byte) error {
	r.cfg = r.decodeDiscoveryRoleConfig(config)
	if !r.cfg.Enabled || extconfig.Get().ListenOnlyMode {
		return nil
	}
	r.startWatchSubnets()
	return nil
}

func (r *DiscoveryRole) Stop() {
}

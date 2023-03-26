package api

import (
	"context"
	"fmt"
	"strings"

	"beryju.io/gravity/pkg/extconfig"
	"beryju.io/gravity/pkg/instance/types"
	"beryju.io/gravity/pkg/roles"
	"beryju.io/gravity/pkg/roles/backup"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type APIMember struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}
type APIMembersOutput struct {
	Members []APIMember `json:"members"`
}

func (r *Role) APIClusterMembers() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *APIMembersOutput) error {
		members, err := r.i.KV().MemberList(ctx)
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		for _, mem := range members.Members {
			output.Members = append(output.Members, APIMember{
				ID:   mem.ID,
				Name: mem.Name,
			})
		}
		return nil
	})
	u.SetName("api.get_members")
	u.SetTitle("Etcd members")
	u.SetTags("roles/api")
	u.SetExpectedErrors(status.Internal)
	return u
}

type APIMemberJoinInput struct {
	Peer       string `json:"peer" maxLength:"255"`
	Roles      string `json:"roles"`
	Identifier string `json:"identifier"`
}
type APIMemberJoinOutput struct {
	EtcdInitialCluster string `json:"etcdInitialCluster"`
}

func (r *Role) APIClusterJoin() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input APIMemberJoinInput, output *APIMemberJoinOutput) error {
		r.i.DispatchEvent(backup.EventTopicBackupRun, roles.NewEvent(ctx, map[string]interface{}{}))
		initialCluster := []string{}
		members, err := r.i.KV().MemberList(ctx)
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		for _, mem := range members.Members {
			for _, u := range mem.PeerURLs {
				initialCluster = append(initialCluster, fmt.Sprintf("%s=%s", mem.Name, u))
			}
		}
		initialCluster = append(initialCluster, fmt.Sprintf(
			"%s=%s", input.Identifier, input.Peer,
		))

		// Pre-configure roles for new node
		roles := strings.Split(input.Roles, ",")
		if input.Roles == "" {
			roles = strings.Split(extconfig.Get().BootstrapRoles, ",")
			// If we're copying our roles, exclude backup
			for idx, role := range roles {
				if role == "backup" {
					slices.Delete(roles, idx, idx+1)
				}
			}
		}
		r.i.KV().Put(
			ctx,
			r.i.KV().Key(
				types.KeyInstance,
				input.Identifier,
				types.KeyRoles,
			).String(),
			strings.Join(roles, ","),
		)

		go func() {
			_, err = r.i.KV().MemberAdd(ctx, []string{input.Peer})
			if err != nil {
				r.log.Warn("failed to add member", zap.Error(err))
			}
		}()

		output.EtcdInitialCluster = strings.Join(initialCluster, ",")
		return nil
	})
	u.SetName("etcd.join_member")
	u.SetTitle("Etcd join")
	u.SetTags("roles/etcd")
	u.SetExpectedErrors(status.Internal)
	return u
}

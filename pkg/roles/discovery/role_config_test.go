package discovery_test

import (
	"testing"

	"beryju.io/gravity/pkg/instance"
	"beryju.io/gravity/pkg/roles/discovery"
	"beryju.io/gravity/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestAPIRoleConfigGet(t *testing.T) {
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("discovery", ctx)
	role := discovery.New(inst)
	role.Start(ctx, []byte{})
	defer role.Stop()

	var output discovery.APIRoleConfigOutput
	assert.NoError(t, role.APIRoleConfigGet().Interact(ctx, struct{}{}, &output))
	assert.NotNil(t, output)
}

func TestAPIRoleConfigPut(t *testing.T) {
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("discovery", ctx)
	role := discovery.New(inst)
	role.Start(ctx, []byte{})
	defer role.Stop()

	assert.NoError(t, role.APIRoleConfigPut().Interact(ctx, discovery.APIRoleConfigInput{
		Config: discovery.RoleConfig{},
	}, &struct{}{}))
}

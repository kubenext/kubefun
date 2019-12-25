package action_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/action/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	alerter := fake.NewMockAlerter(controller)
	logger := log.NopLogger()

	m := action.NewManager(logger)
	payloadRan := false
	fn := func(context.Context, action.Alerter, action.Payload) error {
		payloadRan = true
		return nil
	}

	actionPath := "path"
	err := m.Register(actionPath, fn)
	require.NoError(t, err)

	payload := action.Payload{}
	ctx := context.Background()

	err = m.Dispatch(ctx, alerter, actionPath, payload)
	require.NoError(t, err)
	assert.True(t, payloadRan)

}

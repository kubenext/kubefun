package cluster

import (
	"context"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestFromKubeConfig(t *testing.T) {
	kubeConfig := filepath.Join("testdata", "kubeconfig.yaml")
	config := RESTConfigOptions{}

	_, err := FromKubeConfig(context.TODO(), kubeConfig, "", "", config)
	require.NoError(t, err)
}

package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_bindViper_KUBECONFIG(t *testing.T) {
	cmd := &cobra.Command{}

	expected := "/testdata/kubeconfig.yml"
	os.Setenv("KUBECONFIG", expected)
	defer os.Unsetenv("KUBECONFIG")

	// Before bindViper
	actual := viper.GetString("kubeconfig")
	assert.Equal(t, "", actual)

	err := bindViper(cmd)
	require.NoError(t, err)

	// After bindViper
	actual = viper.GetString("kubeconfig")
	assert.Equal(t, expected, actual)
}

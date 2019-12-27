package dash

import (
	"github.com/kubenext/kubefun/internal/module"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/plugin"
	"github.com/kubenext/kubefun/pkg/plugin/api"
	"github.com/pkg/errors"
)

func initPlugin(moduleManager module.ManagerInterface, actionManager *action.Manager, service api.Service) (*plugin.Manager, error) {
	apiService, err := api.New(service)
	if err != nil {
		return nil, errors.Wrap(err, "create dashboard api")
	}

	m := plugin.NewManager(apiService, moduleManager, actionManager)

	pluginList, err := plugin.AvailablePlugins(plugin.DefaultConfig)
	if err != nil {
		return nil, errors.Wrap(err, "finding available plugins")
	}

	for _, pluginPath := range pluginList {
		if err := m.Load(pluginPath); err != nil {
			return nil, errors.Wrapf(err, "initialize plugin %q", pluginPath)
		}

	}

	return m, nil
}

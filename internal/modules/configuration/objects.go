package configuration

import "github.com/kubenext/kubefun/internal/describer"

var (
	pluginDescriber = &PluginListDescriber{}

	rootDescriber = describer.NewSection(
		"/",
		"Configuration",
		pluginDescriber,
	)
)

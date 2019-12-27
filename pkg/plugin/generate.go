package plugin

//go:generate mockgen -destination=./fake/fakes.go -package=fake github.com/kubenext/kubefun/pkg/plugin Runners,ManagerStore,ClientFactory,ModuleService,Service,Broker
//go:generate mockgen -source=dashboard/dashboard.pb.go -destination=./fake/mock_plugin_client.go -package=fake github.com/kubenext/kubefun/pkg/plugin/dashboard PluginClient
//go:generate mockgen -source=../../vendor/github.com/hashicorp/go-plugin/protocol.go -destination=./fake/mock_client_protocol.go -package=fake github.com/hashicorp/go-plugin ClientProtocol

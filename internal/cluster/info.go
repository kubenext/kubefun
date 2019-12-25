package cluster

import "k8s.io/client-go/tools/clientcmd"

//go:generate mockgen -source=info.go -destination=./fake/mock_info_interface.go -package=fake github.com/kubenext/kubefun/internal/cluster InfoInterface

// InfoInterface provides connection details for a cluster
type InfoInterface interface {
	Context() string
	Cluster() string
	Server() string
	User() string
}

type clusterInfo struct {
	clientConfig clientcmd.ClientConfig
}

func newClusterInfo(clientConfig clientcmd.ClientConfig) clusterInfo {
	return clusterInfo{clientConfig: clientConfig}
}

func (c clusterInfo) Context() string {
	raw, err := c.clientConfig.RawConfig()
	if err != nil {
		return ""
	}
	return raw.CurrentContext
}

func (c clusterInfo) Cluster() string {
	raw, err := c.clientConfig.RawConfig()
	if err != nil {
		return ""
	}

	ktx, ok := raw.Contexts[raw.CurrentContext]
	if !ok || ktx == nil {
		return ""
	}
	return ktx.Cluster
}

func (c clusterInfo) Server() string {
	raw, err := c.clientConfig.RawConfig()
	if err != nil {
		return ""
	}
	ktx, ok := raw.Contexts[raw.CurrentContext]
	if !ok || ktx == nil {
		return ""
	}
	cluster, ok := raw.Clusters[ktx.Cluster]
	if !ok || cluster == nil {
		return ""
	}
	return cluster.Server
}

func (c clusterInfo) User() string {
	raw, err := c.clientConfig.RawConfig()
	if err != nil {
		return ""
	}
	ktx, ok := raw.Contexts[raw.CurrentContext]
	if !ok || ktx == nil {
		return ""
	}
	return ktx.AuthInfo
}

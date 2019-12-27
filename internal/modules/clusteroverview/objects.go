package clusteroverview

import (
	"github.com/kubenext/kubefun/internal/describer"
	"github.com/kubenext/kubefun/pkg/icon"
	"github.com/kubenext/kubefun/pkg/store"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

var (
	customResourcesDescriber = describer.NewCRDSection(
		"/custom-resources",
		"Custom Resources",
	)

	rbacClusterRoles = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/cluster-roles",
		ObjectStoreKey: store.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ListType:       &rbacv1.ClusterRoleList{},
		ObjectType:     &rbacv1.ClusterRole{},
		Titles:         describer.ResourceTitle{List: "RBAC / Cluster Roles", Object: "Cluster Role"},
		ClusterWide:    true,
		IconName:       icon.ClusterOverviewClusterRole,
	})

	rbacClusterRoleBindings = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/cluster-role-bindings",
		ObjectStoreKey: store.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ListType:       &rbacv1.ClusterRoleBindingList{},
		ObjectType:     &rbacv1.ClusterRoleBinding{},
		Titles:         describer.ResourceTitle{List: "RBAC / Cluster Role Bindings", Object: "Cluster Role Binding"},
		ClusterWide:    true,
		IconName:       icon.ClusterOverviewClusterRoleBinding,
	})

	rbacDescriber = describer.NewSection(
		"/rbac",
		"RBAC",
		rbacClusterRoles,
		rbacClusterRoleBindings,
	)

	nodesDescriber = describer.NewResource(describer.ResourceOptions{
		Path:                  "/nodes",
		ObjectStoreKey:        store.Key{APIVersion: "v1", Kind: "Node"},
		ListType:              &v1.NodeList{},
		ObjectType:            &v1.Node{},
		Titles:                describer.ResourceTitle{List: "Nodes", Object: "Node"},
		DisableResourceViewer: true,
		ClusterWide:           true,
		IconName:              icon.ClusterOverviewNode,
	})

	namespacesDescriber = describer.NewResource(describer.ResourceOptions{
		Path:                  "/namespaces",
		ObjectStoreKey:        store.Key{APIVersion: "v1", Kind: "Namespace"},
		ListType:              &v1.NamespaceList{},
		ObjectType:            &v1.Namespace{},
		Titles:                describer.ResourceTitle{List: "Namespaces", Object: "Namespace"},
		DisableResourceViewer: true,
		ClusterWide:           true,
		IconName:              icon.ClusterOverviewNamespace,
	})

	portForwardDescriber = NewPortForwardListDescriber()

	rootDescriber = describer.NewSection(
		"/",
		"Cluster Overview",
		namespacesDescriber,
		customResourcesDescriber,
		rbacDescriber,
		nodesDescriber,
		portForwardDescriber,
	)
)

module github.com/kubenext/kubefun

go 1.13

require (
	github.com/GeertJohan/go.rice v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang/mock v1.3.1
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.2.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/golang-lru v0.5.1
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
	go.opencensus.io v0.22.1
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20191016225839-816a9b7df678
	k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476
	k8s.io/apimachinery v0.0.0-20191016225534-b1267f8c42b4
	k8s.io/client-go v0.0.0-20191016230210-14c42cd304d9
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.14.0
	k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

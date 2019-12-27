module github.com/kubenext/kubefun

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	github.com/GeertJohan/go.rice v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/gobwas/glob v0.2.3
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.2.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/golang-lru v0.5.1
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/magiconair/properties v1.8.1
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/skratchdot/open-golang v0.0.0-20190402232053-79abb63cd66e
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.3.0
	go.opencensus.io v0.22.1
	go.uber.org/zap v1.10.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.21.0
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

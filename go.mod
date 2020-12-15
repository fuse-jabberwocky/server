module github.com/fusejw/server

go 1.15

require (
	github.com/apache/camel-k/pkg/apis/camel v0.0.0
	github.com/apache/camel-k/pkg/client/camel v0.0.0
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee // indirect
	github.com/docker/cli v0.0.0-20200210162036-a4bedce16568 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/go-logr/logr v0.3.0 // indirect
	github.com/go-logr/zapr v0.2.0 // indirect
	github.com/google/go-containerregistry v0.1.4 // indirect
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/rs/dnscache v0.0.0-20190621150935-06bb5526f76b // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/tsenart/go-tsz v0.0.0-20180814235614-0bd30b3df1c3 // indirect
	github.com/tsenart/vegeta v12.7.1-0.20190725001342-b5f4fca92137+incompatible // indirect
	go.uber.org/goleak v1.1.10 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.19.2 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	knative.dev/serving v0.19.0
	sigs.k8s.io/controller-runtime v0.6.4
	sigs.k8s.io/controller-tools v0.0.0-20200528125929-5c0c6ae3b64b // indirect
)

replace (
	github.com/apache/camel-k/pkg/apis/camel => github.com/apache/camel-k/pkg/apis/camel v0.0.0-20201030155835-ccf1c2755f2a
	github.com/apache/camel-k/pkg/client/camel => github.com/apache/camel-k/pkg/client/camel v0.0.0-20201030155835-ccf1c2755f2a
	k8s.io/client-go => k8s.io/client-go v0.18.9
	k8s.io/code-generator => k8s.io/code-generator v0.18.9
	knative.dev/serving => knative.dev/serving v0.18.0
)

// Local modules
replace (
	github.com/fusejw/server/internal/pkg/strimzi => ./internal/pkg/strimzi
	github.com/fusejw/server/internal/server => ./internal/server
)

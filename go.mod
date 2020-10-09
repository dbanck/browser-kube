module github.com/dbanck/browser-kube

go 1.13

require (
	github.com/envoyproxy/go-control-plane v0.6.9 // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible
	github.com/gogo/googleapis v1.1.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/lyft/protoc-gen-validate v0.0.13 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.7
	github.com/spf13/pflag v1.0.5
	github.com/virtual-kubelet/node-cli v0.1.3-0.20200630204823-dccbe1399f29
	github.com/virtual-kubelet/tensile-kube v0.0.0-20200925030435-a9f7766c9cd5
	github.com/virtual-kubelet/virtual-kubelet v1.2.1-0.20200629225228-bd977cb22454
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.18.4
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/component-base v0.17.0
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20200410163147-594e756bea31 // indirect
	k8s.io/kube-scheduler v0.0.0
	k8s.io/kubernetes v1.17.6
	sigs.k8s.io/descheduler v0.10.1-0.20200508041036-423ee35846a8
	sigs.k8s.io/structured-merge-diff v1.0.1 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.1
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/apiserver => k8s.io/apiserver v0.16.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
	k8s.io/component-base => k8s.io/component-base v0.18.4
	k8s.io/cri-api => k8s.io/cri-api v0.16.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.4
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410163147-594e756bea31
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.4
	k8s.io/kubectl => k8s.io/kubectl v0.16.4
	k8s.io/kubelet => k8s.io/kubelet v0.16.4
	k8s.io/kubernetes => k8s.io/kubernetes v1.16.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.4
	k8s.io/metrics => k8s.io/metrics v0.16.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.4
	sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v1.0.1
)

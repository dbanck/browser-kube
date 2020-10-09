package provider

import (
	"github.com/virtual-kubelet/node-cli/manager"
	"github.com/virtual-kubelet/node-cli/opts"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/tensile-kube/pkg/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
)

// ClientConfig defines the configuration of a lower cluster
type ClientConfig struct {
	// allowed qps of the kube client
	KubeClientQPS int
	// allowed burst of the kube client
	KubeClientBurst int
	// config path of the kube client
	ClientKubeConfigPath string
}

// clientCache wraps the lister of client cluster
type clientCache struct {
	podLister    v1.PodLister
	nsLister     v1.NamespaceLister
	cmLister     v1.ConfigMapLister
	secretLister v1.SecretLister
	nodeLister   v1.NodeLister
}

// VirtualK8S is the key struct to implement the tensile kubernetes
type VirtualK8S struct {
	master               kubernetes.Interface
	client               kubernetes.Interface
	config               *rest.Config
	nodeName             string
	version              string
	daemonPort           int32
	ignoreLabels         []string
	clientCache          clientCache
	rm                   *manager.ResourceManager
	updatedNode          chan *corev1.Node
	updatedPod           chan *corev1.Pod
	enableServiceAccount bool
	stopCh               <-chan struct{}
	providerNode         *common.ProviderNode
	configured           bool
}

// NewVirtualK8S creates a connection to a browser
func NewVirtualK8S(cfg provider.InitConfig, cc *ClientConfig,
	ignoreLabelsStr string, enableServiceAccount bool, opts *opts.Opts) (*VirtualK8S, error) {
	return nil, nil
}

// GetClient return the kube client of browsers
func (v *VirtualK8S) GetClient() kubernetes.Interface {
	return v.client
}

// GetMaster return the kube client of provider cluster
func (v *VirtualK8S) GetMaster() kubernetes.Interface {
	return v.master
}

// GetNameSpaceLister returns the namespace cache
func (v *VirtualK8S) GetNameSpaceLister() v1.NamespaceLister {
	return v.clientCache.nsLister
}

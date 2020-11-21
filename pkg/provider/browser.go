package provider

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/virtual-kubelet/node-cli/manager"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"go.opencensus.io/trace"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

// BrowserProvider implements the virtual-kubelet provider interface
type BrowserProvider struct {
	nodeName           string
	daemonEndpointPort int32
	operatingSystem    string
	internalIP         string
	apiPort            int

	pods map[string]*v1.Pod
}

// NewBrowserProvider creates a new Browser Provider
func NewBrowserProvider(config string, rm *manager.ResourceManager, nodeName, operatingSystem string, internalIP string, daemonEndpointPort int32, clusterDomain string, apiPort int) (*BrowserProvider, error) {
	ctx := context.Background()
	p := BrowserProvider{}
	p.pods = map[string]*v1.Pod{}
	p.operatingSystem = operatingSystem
	p.nodeName = nodeName
	p.internalIP = internalIP
	p.daemonEndpointPort = daemonEndpointPort
	p.apiPort = apiPort

	log.G(ctx).Infof("Starting node name %v serving the API on port %v", nodeName, apiPort)

	return &p, nil
}

func getPodName(pod *v1.Pod) string {
	return strings.Join([]string{pod.Namespace, pod.Name}, "/")
}

// CreatePod takes a Kubernetes Pod and deploys it within the provider.
func (p *BrowserProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.CreatePod")
	defer span.End()
	log.G(ctx).Infof("Creating pod %v", pod.Name)

	p.pods[getPodName(pod)] = pod

	return nil
}

// GetPod retrieves a pod by name from the provider (can be cached).
// The Pod returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *BrowserProvider) GetPod(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "browser.GetPod")
	defer span.End()
	log.G(ctx).Infof("Reading pod %v/%v", namespace, name)

	return nil, errors.New("TODO: Implement fetching pods")
}

// GetPodStatus retrieves the status of a pod by name from the provider.
// The PodStatus returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *BrowserProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	ctx, span := trace.StartSpan(ctx, "browser.GetPodStatus")
	defer span.End()
	log.G(ctx).Infof("Reading pod status %v/%v", namespace, name)

	return nil, errors.New("TODO: implement fetching pod status")
}

// GetPodStats gets the metrics for a pod. As the browser does not provide the metrics needed we stub it out
func (p *BrowserProvider) GetPodStats(ctx context.Context, namespace, name string) *stats.PodStats {
	ctx, span := trace.StartSpan(ctx, "browser.GetPodStats")
	defer span.End()
	log.G(ctx).Infof("Reading pod stats %v/%v", namespace, name)

	podRef := stats.PodReference{Name: name, Namespace: namespace, UID: name}
	return &stats.PodStats{PodRef: podRef}
}

// GetPods retrieves a list of all pods running on the provider (can be cached).
// The Pods returned are expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (p *BrowserProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "browser.GetPods")
	defer span.End()
	log.G(ctx).Infof("Listing pods %+v", p.pods)

	pods := []*v1.Pod{}
	for _, pod := range p.pods {
		pods = append(pods, pod)
	}

	return pods, nil
}

// UpdatePod takes a Kubernetes Pod and updates it within the provider.
func (p *BrowserProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.UpdatePod")
	defer span.End()
	log.G(ctx).Infof("Updating pod %v", pod.Name)

	p.pods[getPodName(pod)] = pod

	return nil
}

// DeletePod takes a Kubernetes Pod and deletes it from the provider. Once a pod is deleted, the provider is
// expected to call the NotifyPods callback with a terminal pod status where all the containers are in a terminal
// state, as well as the pod. DeletePod may be called multiple times for the same pod.
func (p *BrowserProvider) DeletePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.DeletePod")
	defer span.End()
	log.G(ctx).Infof("Deleting pod %v", pod.Name)

	delete(p.pods, getPodName(pod))

	return nil
}

// GetContainerLogs retrieves the logs of a container by name from the provider.
func (p *BrowserProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	ctx, span := trace.StartSpan(ctx, "browser.GetContainerLogs")
	defer span.End()
	log.G(ctx).Infof("Getting container logs for %v/%v %v", namespace, podName, containerName)

	logContent := "TODO: implement"

	return ioutil.NopCloser(strings.NewReader(logContent)), nil
}

// RunInContainer executes a command in a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (p *BrowserProvider) RunInContainer(ctx context.Context, namespace, name, container string, cmd []string, attach api.AttachIO) error {
	log.G(ctx).Infof("Running in container %v/%v %v", namespace, name, container)
	return errors.New("TODO: implement RunInContainer")
}

// ConfigureNode enables a provider to configure the node object that
// will be used for Kubernetes.
func (p *BrowserProvider) ConfigureNode(ctx context.Context, node *v1.Node) {
	log.G(ctx).Infof("Configuring Node")

	capacity := v1.ResourceList{
		v1.ResourceCPU:    resource.MustParse("10000"),
		v1.ResourceMemory: resource.MustParse("4Ti"),
		v1.ResourcePods:   resource.MustParse("5000"),
	}

	node.Status.Capacity = capacity
	node.Status.Allocatable = capacity
	node.Status.Conditions = []v1.NodeCondition{
		{
			Type:               "Ready",
			Status:             v1.ConditionTrue,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletReady",
			Message:            "kubelet is ready.",
		},
		{
			Type:               "OutOfDisk",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientDisk",
			Message:            "kubelet has sufficient disk space available",
		},
		{
			Type:               "MemoryPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientMemory",
			Message:            "kubelet has sufficient memory available",
		},
		{
			Type:               "DiskPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasNoDiskPressure",
			Message:            "kubelet has no disk pressure",
		},
		{
			Type:               "NetworkUnavailable",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "RouteCreated",
			Message:            "RouteController created a route",
		},
	}
	node.Status.Addresses = []v1.NodeAddress{
		{
			Type:    "InternalIP",
			Address: p.internalIP,
		},
	}
	node.Status.DaemonEndpoints = v1.NodeDaemonEndpoints{
		KubeletEndpoint: v1.DaemonEndpoint{
			Port: p.daemonEndpointPort,
		},
	}
	node.Status.NodeInfo.OperatingSystem = p.operatingSystem
	node.ObjectMeta.Labels["alpha.service-controller.kubernetes.io/exclude-balancer"] = "true"
}

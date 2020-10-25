package provider

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/virtual-kubelet/node-cli/manager"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"go.opencensus.io/trace"
	v1 "k8s.io/api/core/v1"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

// BrowserProvider implements the virtual-kubelet provider interface
type BrowserProvider struct {
	nodeName string

	metricsSync     sync.Mutex
	metricsSyncTime time.Time
	lastMetric      *stats.Summary
}

// NewBrowserProvider creates a new Browser Provider
func NewBrowserProvider(config string, rm *manager.ResourceManager, nodeName, operatingSystem string, internalIP string, daemonEndpointPort int32, clusterDomain string) (*BrowserProvider, error) {
	p := BrowserProvider{nodeName: nodeName}
	return &p, nil
}

// CreatePod takes a Kubernetes Pod and deploys it within the provider.
func (p *BrowserProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.CreatePod")
	defer span.End()
	log.G(ctx).Infof("Creating pod %v", pod.Name)

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
	log.G(ctx).Infof("Listing pods")
	log.G(ctx).Errorf("TODO: implement listing pods")

	var pods []*v1.Pod
	return pods, nil
}

// UpdatePod takes a Kubernetes Pod and updates it within the provider.
func (p *BrowserProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.UpdatePod")
	defer span.End()
	log.G(ctx).Infof("Updating pod %v", pod.Name)
	return nil
}

// DeletePod takes a Kubernetes Pod and deletes it from the provider. Once a pod is deleted, the provider is
// expected to call the NotifyPods callback with a terminal pod status where all the containers are in a terminal
// state, as well as the pod. DeletePod may be called multiple times for the same pod.
func (p *BrowserProvider) DeletePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "browser.DeletePod")
	defer span.End()

	log.G(ctx).Infof("Deleting pod %v", pod.Name)
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
}

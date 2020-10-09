/*
 * Copyright ©2020. The virtual-kubelet authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package provider

import (
	"context"
	"fmt"
	"io"

	"github.com/virtual-kubelet/virtual-kubelet/node"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

var _ node.PodLifecycleHandler = &VirtualK8S{}
var _ node.PodNotifier = &VirtualK8S{}
var _ node.NodeProvider = &VirtualK8S{}

// CreatePod takes a Kubernetes Pod and deploys it within the provider.
func (v *VirtualK8S) CreatePod(ctx context.Context, pod *corev1.Pod) error {
	if pod.Namespace == "kube-system" {
		return nil
	}

	klog.V(3).Infof("Creating pod %v/%+v", pod.Namespace, pod.Name)
	return nil
}

// UpdatePod takes a Kubernetes Pod and updates it within the provider.
func (v *VirtualK8S) UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	if pod.Namespace == "kube-system" {
		return nil
	}
	klog.V(3).Infof("Updating pod %v/%+v", pod.Namespace, pod.Name)
	return nil
}

// DeletePod takes a Kubernetes Pod and deletes it from the provider.
func (v *VirtualK8S) DeletePod(ctx context.Context, pod *corev1.Pod) error {
	if pod.Namespace == "kube-system" {
		return nil
	}
	klog.V(3).Infof("Deleting pod %v/%+v", pod.Namespace, pod.Name)

	return nil
}

// GetPod retrieves a pod by name from the provider (can be cached).
// The Pod returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (v *VirtualK8S) GetPod(ctx context.Context, namespace string, name string) (*corev1.Pod, error) {
	return nil, nil
}

// GetPodStatus retrieves the status of a pod by name from the provider.
// The PodStatus returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (v *VirtualK8S) GetPodStatus(ctx context.Context, namespace string, name string) (*corev1.PodStatus, error) {
	return nil, nil
}

// GetPods retrieves a list of all pods running on the provider (can be cached).
// The Pods returned are expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (v *VirtualK8S) GetPods(_ context.Context) ([]*corev1.Pod, error) {
	return nil, nil
}

// GetContainerLogs retrieves the logs of a container by name from the provider.
func (v *VirtualK8S) GetContainerLogs(ctx context.Context, namespace string,
	podName string, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	return nil, nil
}

// RunInContainer executes a command in a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (v *VirtualK8S) RunInContainer(ctx context.Context, namespace string, podName string, containerName string, cmd []string, attach api.AttachIO) error {
	return fmt.Errorf("Not supported")
}

// NotifyPods instructs the notifier to call the passed in function when
// the pod status changes. It should be called when a pod's status changes.
//
// The provided pointer to a Pod is guaranteed to be used in a read-only
// fashion. The provided pod's PodStatus should be up to date when
// this function is called.
//
// NotifyPods will not block callers.
func (v *VirtualK8S) NotifyPods(ctx context.Context, f func(*corev1.Pod)) {
	klog.Info("Called NotifyPods")
}

// createSecrets takes a Kubernetes Pod and deploys it within the provider.
func (v *VirtualK8S) createSecrets(ctx context.Context, secrets []string, ns string) error {
	return nil
}

// createConfigMaps a Kubernetes Pod and deploys it within the provider.
func (v *VirtualK8S) createConfigMaps(ctx context.Context, configmaps []string, ns string) error {
	return nil
}

// deleteConfigMaps a Kubernetes Pod and deploys it within the provider.
func (v *VirtualK8S) deleteConfigMaps(ctx context.Context, configmaps []string, ns string) error {
	return nil
}

// createPVCs a Kubernetes Pod and deploys it within the provider.
func (v *VirtualK8S) createPVCs(ctx context.Context, pvcs []string, ns string) error {
	return nil
}

func (v *VirtualK8S) patchConfigMap(cm, clone *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	return nil, nil
}

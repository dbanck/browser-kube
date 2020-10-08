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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

// ConfigureNode enables a provider to configure the node object that
// will be used for Kubernetes.
func (v *VirtualK8S) ConfigureNode(ctx context.Context, node *corev1.Node) {
	return
}

// Ping tries to connect to client cluster
// implement node.NodeProvider
func (v *VirtualK8S) Ping(ctx context.Context) error {
	// If node or master ping fail, we should it as a failed ping
	_, err := v.master.Discovery().ServerVersion()
	if err != nil {
		klog.Error("Failed ping")
		return fmt.Errorf("could not list master apiserver statuses: %v", err)
	}
	_, err = v.client.Discovery().ServerVersion()
	if err != nil {
		klog.Error("Failed ping")
		return fmt.Errorf("could not list client apiserver statuses: %v", err)
	}
	return nil
}

// NotifyNodeStatus is used to asynchronously monitor the node.
// The passed in callback should be called any time there is a change to the
// node's status.
// This will generally trigger a call to the Kubernetes API server to update
// the status.
//
// NotifyNodeStatus should not block callers.
func (v *VirtualK8S) NotifyNodeStatus(ctx context.Context, f func(*corev1.Node)) {
	klog.Info("Called NotifyNodeStatus")
	go func() {
		for {
			select {
			case node := <-v.updatedNode:
				klog.Infof("Enqueue updated node %v", node.Name)
				f(node)
			case <-v.stopCh:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

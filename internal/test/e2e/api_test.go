// +build e2e

package e2e

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	c *kubernetes.Clientset
	e *httpexpect.Expect
)

func TestApi(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	c = client
	e = httpexpect.New(t, "http://localhost:80")
	t.Run("Pods", testPods)
}

func testPods(t *testing.T) {
	t.Run("empty object with no pods", func(t *testing.T) {
		e.GET("/pods").
			Expect().
			Status(http.StatusOK).JSON().Object()
	})

	t.Run("one object with a scheduled pod", func(t *testing.T) {
		ctx := context.Background()
		req := &core.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-test-pod",
				Namespace: "default",
				Labels: map[string]string{
					"app": "demo",
				},
			},
			Spec: core.PodSpec{
				NodeName: "vkubelet-browser",
				Containers: []core.Container{
					{
						Name:            "busybox",
						Image:           "busybox",
						ImagePullPolicy: core.PullIfNotPresent,
						Command: []string{
							"sleep",
							"3600",
						},
					},
				},
			},
		}

		if _, err := c.CoreV1().Pods("default").Create(ctx, req, metav1.CreateOptions{}); err != nil {
			fmt.Errorf(err.Error())
			t.Fail()
		}

		pods := e.GET("/pods").
			Expect().
			Status(http.StatusOK).JSON()
		pods.Object().ContainsKey("default/my-test-pod")
	})
}

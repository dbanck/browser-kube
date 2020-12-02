// +build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"net/http"

	"github.com/gavv/httpexpect/v2"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	vke2e "github.com/virtual-kubelet/virtual-kubelet/test/e2e"
)

const (
	defaultNamespace = core.NamespaceDefault
	defaultNodeName  = "vkubelet-mock-0"
)

var (
	kubeconfig string
	namespace  string
	nodeName   string
	c          *kubernetes.Clientset
)

// go1.13 compatibility cf. https://github.com/golang/go/issues/31859
var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to the kubeconfig file to use when running the test suite outside a kubernetes cluster")
	flag.StringVar(&namespace, "namespace", defaultNamespace, "the name of the kubernetes namespace to use for running the test suite (i.e. where to create pods)")
	flag.StringVar(&nodeName, "node-name", defaultNodeName, "the name of the virtual-kubelet node to test")
	flag.Parse()
}

// Provider-specific setup function
func setup() error {
	fmt.Println("Setting up end-to-end test suite for browser-kube provider...")
	return nil
}

// Provider-specific teardown function
func teardown() error {
	fmt.Println("Tearing down end-to-end test suite for browser-kube provider...")
	return nil
}

//  Provider-specific shouldSkipTest function
func shouldSkipTest(testName string) bool {
	skippedTests := [...]string{"TestCreatePodWithMandatoryInexistentConfigMap", "TestCreatePodWithMandatoryInexistentSecrets", "TestCreatePodWithOptionalInexistentConfigMap", "TestCreatePodWithOptionalInexistentSecrets", "TestGetPods", "TestGetStatsSummary", "TestNodeCreateAfterDelete", "TestPodLifecycleForceDelete", "TestPodLifecycleGracefulDelete"}

	for _, item := range skippedTests {
		if item == testName {
			return true
		}
	}
	fmt.Println("Executing", testName)
	return false
}

// TestEndToEnd creates and runs the end-to-end test suite for virtual kubelet
func TestEndToEnd(t *testing.T) {
	setDefaults()
	config := vke2e.EndToEndTestSuiteConfig{
		Kubeconfig:     kubeconfig,
		Namespace:      namespace,
		NodeName:       nodeName,
		Setup:          setup,
		Teardown:       teardown,
		ShouldSkipTest: shouldSkipTest,
	}
	fmt.Println("Starting virtual-kubelet E2E tests")
	ts := vke2e.NewEndToEndTestSuite(config)
	ts.Run(t)
}

func wsHttpHandlerTester(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:80",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func TestPodsApi(t *testing.T) {
	t.Run("Pods", func(t *testing.T) {
		ctx := context.Background()
		ns := "pod-test"
		createNamespace(ctx, ns)
		e := wsHttpHandlerTester(t)
		e.GET("/pods").
			Expect().
			Status(http.StatusOK).JSON().Object()

		firstBrowser := e.GET("/ws").WithWebsocketUpgrade().
			Expect().
			Status(http.StatusSwitchingProtocols).
			Websocket()
		defer firstBrowser.Disconnect()

		err := schedulePodWithName(ctx, ns, "my-test-pod")
		if err != nil {
			t.Errorf("Could not schedule pod: %s", err.Error())
		}

		firstBrowser.Expect().TextMessage().Body().Contains("my-test-pod")
		e.GET("/pods").
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsKey(fmt.Sprintf("%s/my-test-pod", ns))

		err = schedulePodWithName(ctx, ns, "my-other-test-pod")
		if err != nil {
			t.Errorf("Could not schedule pod: %s", err.Error())
		}
		firstBrowser.Expect().TextMessage().Body().Contains("my-other-test-pod")

		secondBrowser := e.GET("/ws").WithWebsocketUpgrade().
			Expect().
			Status(http.StatusSwitchingProtocols).
			Websocket()
		defer secondBrowser.Disconnect()

		// Two pods in the ns, order can be random
		secondBrowser.Expect().TextMessage().Body().Contains(ns)
		secondBrowser.Expect().TextMessage().Body().Contains(ns)

		deletePod(ctx, ns, "my-test-pod")
		deletePod(ctx, ns, "my-other-test-pod")
		deleteNamespace(ctx, ns)
	})
}

// setDefaults sets sane defaults in case no values (or empty ones) have been provided.
func setDefaults() {
	if namespace == "" {
		namespace = defaultNamespace
	}
	if nodeName == "" {
		nodeName = defaultNodeName
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	c = client
}

func createNamespace(ctx context.Context, ns string) {
	nsSpec := &core.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}

	_, err := c.CoreV1().Namespaces().Create(ctx, nsSpec, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}

func deleteNamespace(ctx context.Context, namespace string) {
	gracePeriod := int64(0)
	policy := metav1.DeletePropagationOrphan
	c.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{GracePeriodSeconds: &gracePeriod, PropagationPolicy: &policy})
	time.Sleep(5 * time.Second)
}

func schedulePodWithName(ctx context.Context, namespace string, name string) error {
	req := &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
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

	_, err := c.CoreV1().Pods(namespace).Create(ctx, req, metav1.CreateOptions{})
	return err
}

func deletePod(ctx context.Context, namespace string, name string) {
	c.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

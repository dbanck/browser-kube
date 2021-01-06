package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	browserProvider "github.com/dbanck/browser-kube/pkg/provider"
	"github.com/sirupsen/logrus"
	cli "github.com/virtual-kubelet/node-cli"
	logruscli "github.com/virtual-kubelet/node-cli/logrus"
	"github.com/virtual-kubelet/node-cli/opts"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	logruslogger "github.com/virtual-kubelet/virtual-kubelet/log/logrus"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"github.com/virtual-kubelet/virtual-kubelet/trace/opencensus"
)

var (
	buildVersion = "N/A"
	buildTime    = "N/A"
	k8sVersion   = "v1.19.2" // This should follow the version of k8s.io/kubernetes we are importing
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	logger := logrus.StandardLogger()
	log.L = logruslogger.FromLogrus(logrus.NewEntry(logger))
	logConfig := &logruscli.Config{LogLevel: "info"}

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.G(ctx).Fatal(err)
	}

	trace.T = opencensus.Adapter{}

	o, err := opts.FromEnv()
	if err != nil {
		log.G(ctx).Fatal(err)
	}
	o.Provider = "browser"
	o.Version = strings.Join([]string{k8sVersion, buildVersion}, "-")

	node, err := cli.New(ctx,
		cli.WithBaseOpts(o),
		cli.WithCLIVersion(buildVersion, buildTime),
		cli.WithProvider("browser", func(cfg provider.InitConfig) (provider.Provider, error) {
			return browserProvider.NewBrowserProvider(cfg.ConfigPath, cfg.ResourceManager, cfg.NodeName, cfg.OperatingSystem, cfg.InternalIP, cfg.DaemonPort, cfg.KubeClusterDomain, apiPort)
		}),
		cli.WithPersistentFlags(logConfig.FlagSet()),
		cli.WithPersistentPreRunCallback(func() error {
			return logruscli.Configure(logConfig, logger)
		}),
	)

	if err != nil {
		log.G(ctx).Fatal(err)
	}

	if err := node.Run(ctx); err != nil {
		log.G(ctx).Fatal(err)
	}
}

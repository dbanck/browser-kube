# Contributing

## Running the E2E test

Please install [skaffold](https://skaffold.dev/) and [kind](https://kind.sigs.k8s.io/).
Run `make kindE2E` to run the e2e tests in a freshly created kind cluster.
Run `export KUBECONFIG="/path/to/project/root/kubeconfig.yaml"` to connect your kubectl to the kind e2e cluster for debugging. The Makefile ensures that the cluster is reused, so if you need a new cluster please run `make kindE2E.clean`.

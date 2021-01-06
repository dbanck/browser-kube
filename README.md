# browser-kube

## Contributing

### Prerequisites

- We need Dockerhub credentials in environment variables (`DOCKER_PASSWORD` & `DOCKER_USERNAME`)
- kubectl, docker, go, skaffold, and kind should be installed

### Development

Run `make dev` to start a development environment (cluster using kind, deployment using skaffold) that automatically reloads on changes

### Testing

Run `make e2e` for our e2e tests & ensure new functionality you provide is covered in an E2E test.

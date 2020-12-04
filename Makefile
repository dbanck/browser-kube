GIT_COMMIT=$(shell git rev-parse "HEAD^{commit}")
VERSION=$(shell git describe --tags --abbrev=14 "${GIT_COMMIT}^{commit}" --always)
BUILD_TIME=$(shell TZ=Asia/Shanghai date +%FT%T%z)
binary := virtual-kubelet

include Makefile.cluster
include Makefile.e2e
include Makefile.dev

.PHONY: all
all: test build

.PHONY: build
build: fmt vet provider

.PHONY: fmt
fmt:
	go fmt ./pkg/...

.PHONY: vet
vet:
	go vet ./pkg/...

.PHONY: provider
provider: OUTPUT_DIR ?= bin
provider:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'main.buildVersion=$(VERSION)' -X 'main.buildTime=${BUILD_TIME}'" -o $(OUTPUT_DIR)/$(binary) ./cmd/provider

.PHONY: test
test:
	go test -v -count=1 ./...
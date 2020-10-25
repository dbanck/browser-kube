GIT_COMMIT=$(shell git rev-parse "HEAD^{commit}")
VERSION=$(shell git describe --tags --abbrev=14 "${GIT_COMMIT}^{commit}" --always)
BUILD_TIME=$(shell TZ=Asia/Shanghai date +%FT%T%z)
binary := virtual-kubelet


include Makefile.e2e
include Makefile.dev

all: test build

build: fmt vet provider

fmt:
	go fmt ./pkg/...

vet:
	go vet ./pkg/...

provider: OUTPUT_DIR ?= bin
provider:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'main.buildVersion=$(VERSION)' -X 'main.buildTime=${BUILD_TIME}'" -o $(OUTPUT_DIR)/$(binary) ./cmd/provider

test:
	go test -count=1 ./pkg/...
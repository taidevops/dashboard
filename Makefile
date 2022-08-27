SHELL = /bin/bash
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)

CODEGEN_VERSION := v0.24.1
CODEGEN_BIN := $(GOPATH)/pkg/mod/k8s.io/code-generator@$(CODEGEN_VERSION)/generate-groups.sh

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MIN_GO_MAJOR_VERSION = 1
MIN_GO_MINOR_VERSION = 19
GO_BINARY := $(shell which go)

MAIN_PACKAGE = github.com/taidevops/dashboard/src/app/backend
KUBECONFIG ?= $(HOME)/.kube/config
SERVE_BINARY = .tmp/serve/dashboard
RELEASE_VERSION = v2.6.1

.PHONY: ensure-version
ensure-version:
	node ./aio/scripts/version.mjs

.PHONY: ensure-codegen
ensure-codegen: ensure-go
	go get -d k8s.io/code-generator@$(CODEGEN_VERSION)
	go mod tidy
	chmod +x $(CODEGEN_BIN)

.PHONY: ensure-go
ensure-go:
ifndef GO_BINARY
	$(error "Cannot find go binary")
endif
	@if [ $(GO_MAJOR_VERSION) -gt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 0 ;\
  elif [ $(GO_MAJOR_VERSION) -lt $(MIN_GO_MAJOR_VERSION) ]; then \
		exit 1; \
  elif [ $(GO_MINOR_VERSION) -lt $(MIN_GO_MINOR_VERSION) ] ; then \
		exit 1; \
  fi

.PHONY: clean
clean:
	rm -rf .tmp

.PHONY: build-backend
build-backend: ensure-go
	CGO_ENABLED=0 go build -ldflags "-X $(MAIN_PACKAGE)/client.Version=$(RELEASE_VERSION)" -gcflags="all=-N -l" -o $(SERVE_BINARY) $(MAIN_PACKAGE)

.PHONY: build
build: clean ensure-go
	./aio/scripts/build.sh

.PHONY: serve-backend
serve-backend: build-backend run-backend

.PHONY: run-backend
run-backend:
	$(SERVE_BINARY)

.PHONY: check-codegen
check-codegen: ensure-codegen
	./aio/scripts/verify-codegen.sh

.PHONY: start-cluster
start-cluster:
	./aio/scripts/start-cluster.sh

.PHONY: stop-cluster
stop-cluster:
	./aio/scripts/stop-cluster.sh

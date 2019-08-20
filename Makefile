THIS_DIR  := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
THIS_FILE := $(lastword $(MAKEFILE_LIST))

#------------------------------------------------------------------
# Project build information
#------------------------------------------------------------------
PROJNAME     := version-compliance-checker
VENDOR       := swade1987
MAINTAINER   := steven@stevenwade.co.uk

GIT_REPO    := github.com/$(VENDOR)/$(PROJNAME)
GIT_SHA     := $(shell git rev-parse --verify HEAD)
BUILD_DATE  := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
VERSION     := $(shell cat VERSION)

#------------------------------------------------------------------
# Go configuration
#------------------------------------------------------------------
GOCMD          := go
BIN            := bin
VERSION_PKG    := $(GIT_REPO)/pkg/runtime/version
LD_FLAGS       := -ldflags '-X "$(VERSION_PKG).GitSHA=${GIT_SHA}" -X "$(VERSION_PKG).Version=${VERSION}" -X "$(VERSION_PKG).BuildDate=${BUILD_DATE}" -X "$(VERSION_PKG).Maintainer=${MAINTAINER}"'

# If no target is defined, assume the host is the target.
ifeq ($(origin GOOS), undefined)
	GOOS := $(shell go env GOOS)
endif
# Lots of these target goarches probably won't work,
# since we depend on vendored packages also being built for the correct arch
ifeq ($(origin GOARCH), undefined)
	GOARCH := $(shell go env GOARCH)
endif

TARGETS := darwin linux

#------------------------------------------------------------------
# Build targets
#------------------------------------------------------------------


.PHONY: all
all: test version-compliance-checker  ## Run test and version-compliance-checker  (default)

.PHONY: test
test: fmt vet ## Run tests
	$(GOCMD) test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: version-compliance-checker
version-compliance-checker: fmt vet ## Build londonbikes binary
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build $(LD_FLAGS) -o $(BIN)/$(PROJNAME)-$(GOOS) $(GIT_REPO)

.PHONY: fmt
fmt: ## Run go fmt against code
	$(GOCMD) fmt ./pkg/... ./cmd/...

.PHONY: vet
vet: ## Run go vet against code
	$(GOCMD) vet ./pkg/... ./cmd/...

.PHONY: run
run: ## Run the app
	IOS_REQUIRED_VERSION=1.1.1 ANDROID_REQUIRED_VERSION=2.2.2 $(BIN)/$(PROJNAME)-$(GOOS)

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
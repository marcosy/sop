# Operative System detection
os=$(shell uname -s)
ifeq ($(os),Darwin)
os=darwin
else ifeq ($(os),Linux)
os=linux
else
$(error unsupported operative system: $(os))
endif

# Architecture detection
arch=$(shell uname -m)
ifeq ($(arch),x86_64)
arch=amd64
else ifeq ($(arch),arm64)
arch=amd64
else ifeq ($(arch),aarch64)
arch=arm64
else
$(error unsupported architecture: $(arch))
endif

# Build settings
DIR := ${CURDIR}
build_dir := $(DIR)/.build/${os}-${arch}

# Golang settings
go_version = 1.19.2
go_dir = $(build_dir)/go/$(go_version)
go_url = https://storage.googleapis.com/golang/go$(go_version).$(os)-$(arch).tar.gz
go_path := PATH="$(go_dir)/bin:$(PATH)"

# Golang CI settings
golangci_lint_version = v1.50.1
golangci_lint_dir = $(build_dir)/golangci_lint/$(golangci_lint_version)
golangci_lint_bin = $(golangci_lint_dir)/golangci-lint
golangci_lint_cache = $(golangci_lint_dir)/cache
golangci_lint_url = https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

# Help message settings 
blue := $(shell which tput > /dev/null && tput setaf 4 2>/dev/null || echo "")
reset := $(shell which tput > /dev/null && tput sgr0 2>/dev/null || echo "")
bold  := $(shell which tput > /dev/null && tput bold 2>/dev/null || echo "")
target_max_char=25

# Makefile targets
default: help

##@ Environment setup
go-check: ## Checks (and updates if necesary) the go runtime used for this project
ifneq (go$(go_version), $(shell $(go_path) go version 2>/dev/null | cut -f3 -d' '))
	@echo "Installing go$(go_version)..."
	rm -rf $(dir $(go_dir))
	mkdir -p $(go_dir)
	curl -sSfL $(go_url) | tar xz -C $(go_dir) --strip-components=1
endif

install-golangci-lint: $(golangci_lint_bin) ## Installs golangci lint
$(golangci_lint_bin):
	@echo "Installing golangci-lint $(golangci_lint_version)..."
	rm -rf $(dir $(golangci_lint_dir))
	mkdir -p $(golangci_lint_dir)
	mkdir -p $(golangci_lint_cache)
	curl -sSfL $(golangci_lint_url) | sh -s -- -b $(golangci_lint_dir) $(golangci_lint_version)

##@ Building
.PHONY: all
all: build lint unit-test component-test ## Builds sop binary, lints the code and run tests

.PHONY: build
build: go-check ## Builds sop binary
	@$(go_path) go build -o ./bin/sop ./main.go

.PHONY: unit-test
unit-test: go-check ## Runs unit tests
	@$(go_path) go test ./... --race

.PHONY: component-test
component-test: build ## Run component tests
	@cd test/component && ./run.sh

.PHONY: lint
lint: install-golangci-lint ## Runs golang-ci linter
	@GOLANGCI_LINT_CACHE="$(golangci_lint_cache)" $(golangci_lint_bin) run ./...

##@ Others
.PHONY: help
help: ## Show this help message.
	@awk 'BEGIN {FS = ":.*##"; printf "\n$(bold)Usage:$(reset) make $(blue)<target>$(reset)\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(blue)%-$(target_max_char)s$(reset) %s\n", $$1, $$2 } /^##@/ { printf "\n $(bold)%s$(reset) \n", substr($$0, 5) } ' $(MAKEFILE_LIST) 

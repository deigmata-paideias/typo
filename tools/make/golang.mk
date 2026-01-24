VERSION_PACKAGE := github.com/deigmata-paideias/typo/internal/cmd/version

GO_LDFLAGS += -X $(VERSION_PACKAGE).typoVersion=$(shell cat VERSION) \
	-X $(VERSION_PACKAGE).gitCommitID=$(GIT_COMMIT) \
	-w -s

GIT_COMMIT:=$(shell git rev-parse HEAD)

##@ Golang

.PHONY: fmt
fmt: ## Golang fmt
	go fmt ./...

.PHONY: vet
vet: ## Golang vet
	go vet ./...

.PHONY: dev-run
dev-run: ## Golang dev, run main by run.
	go run cmd/main.go

.PHONY: prod-run
prod-run: ## Golang prod, run bin by run.
	bin/darwin/arm64/collector

# 默认使用当前系统平台，可以通过参数覆盖：make build GOOS=linux GOARCH=amd64
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: build
# build
build: ## Golang build, support cross-compile: make build GOOS=linux GOARCH=amd64
	@version=$$(cat VERSION); \
	echo "Building for $(GOOS)/$(GOARCH)..."; \
	mkdir -p bin/$(GOOS)/$(GOARCH); \
	CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(GOOS)/$(GOARCH)/typo -ldflags "$(GO_LDFLAGS)" cmd/main.go

.PHONY: all-platform-build
all-platform-build: ## Build for all platforms (linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64)
	@echo "Building for all platforms..."
	@$(MAKE) build GOOS=linux GOARCH=amd64
	@$(MAKE) build GOOS=linux GOARCH=arm64
	@$(MAKE) build GOOS=darwin GOARCH=amd64
	@$(MAKE) build GOOS=darwin GOARCH=arm64
	@$(MAKE) build GOOS=windows GOARCH=amd64
	@echo "All platform builds completed. Binaries in bin/"
	@find bin -type f -name typo | xargs ls -lh

.PHONY: go-lint
go-lint: ## run golang lint
	golangci-lint run --config tools/linter/golang-ci/.golangci.yml

.PHONY: test
test: ## run golang test
	go test -v ./...

.PHONY: golang-all
golang-all: ## run fmt lint vet build api test
golang-all: fmt go-lint vet build test

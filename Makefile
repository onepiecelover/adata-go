# AData-Go Makefile

# 变量定义
APP_NAME := adata-go
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION := $(shell go version | awk '{print $$3}')
GIT_COMMIT := $(shell git rev-parse HEAD)

# 构建标志
LDFLAGS := -ldflags "-X github.com/onepiecelover/adata-go.Version=$(VERSION) \
                     -X github.com/onepiecelover/adata-go.BuildTime=$(BUILD_TIME) \
                     -X github.com/onepiecelover/adata-go.GoVersion=$(GO_VERSION) \
                     -X github.com/onepiecelover/adata-go.GitCommit=$(GIT_COMMIT)"

# 默认目标
.PHONY: all
all: clean deps lint test build

# 安装依赖
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# 代码格式化
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 代码检查
.PHONY: lint
lint: fmt
	@echo "Running linter..."
	golangci-lint run ./...

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# 运行测试并生成覆盖率报告
.PHONY: test-coverage
test-coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 运行基准测试
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# 构建二进制文件
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	go build $(LDFLAGS) -o bin/$(APP_NAME) ./examples/basic

# 构建所有示例
.PHONY: build-examples
build-examples:
	@echo "Building examples..."
	mkdir -p bin/examples
	go build $(LDFLAGS) -o bin/examples/basic ./examples/basic
	go build $(LDFLAGS) -o bin/examples/concurrent ./examples/concurrent
	go build $(LDFLAGS) -o bin/examples/stock_info ./examples/stock_info

# 交叉编译
.PHONY: build-cross
build-cross:
	@echo "Cross-compiling..."
	mkdir -p bin/cross
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/cross/$(APP_NAME)-linux-amd64 ./examples/basic
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/cross/$(APP_NAME)-linux-arm64 ./examples/basic
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/cross/$(APP_NAME)-windows-amd64.exe ./examples/basic
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/cross/$(APP_NAME)-darwin-amd64 ./examples/basic
	
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/cross/$(APP_NAME)-darwin-arm64 ./examples/basic

# 运行示例
.PHONY: run-basic
run-basic: build
	@echo "Running basic example..."
	./bin/$(APP_NAME)

.PHONY: run-concurrent
run-concurrent:
	@echo "Running concurrent example..."
	go run ./examples/concurrent

.PHONY: run-stock-info
run-stock-info:
	@echo "Running stock info example..."
	go run ./examples/stock_info

# 清理
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean -cache

# 安装工具
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 生成文档
.PHONY: docs
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# 版本信息
.PHONY: version
version:
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"

# Docker构建
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .
	docker tag $(APP_NAME):$(VERSION) $(APP_NAME):latest

# Docker运行
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run --rm -it $(APP_NAME):latest

# 发布准备
.PHONY: release
release: clean install-tools lint test build-cross
	@echo "Release preparation completed."
	@echo "Binaries are available in bin/cross/"

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all          - Run clean, deps, lint, test, and build"
	@echo "  deps         - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  bench        - Run benchmarks"
	@echo "  build        - Build basic example binary"
	@echo "  build-examples - Build all example binaries"
	@echo "  build-cross  - Cross-compile for multiple platforms"
	@echo "  run-basic    - Run basic example"
	@echo "  run-concurrent - Run concurrent example"
	@echo "  run-stock-info - Run stock info example"
	@echo "  clean        - Clean build artifacts"
	@echo "  install-tools- Install development tools"
	@echo "  docs         - Start documentation server"
	@echo "  version      - Show version information"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  release      - Prepare for release"
	@echo "  help         - Show this help message"
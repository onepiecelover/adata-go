#!/bin/bash

# AData-Go 快速启动脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 输出函数
info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查Go环境
check_go() {
    if ! command -v go &> /dev/null; then
        error "Go is not installed. Please install Go 1.18 or later."
        exit 1
    fi
    
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    info "Go version: $GO_VERSION"
}

# 检查依赖
check_dependencies() {
    info "Checking dependencies..."
    go mod download
    go mod tidy
    success "Dependencies checked"
}

# 运行测试
run_tests() {
    info "Running tests..."
    if go test -v ./...; then
        success "All tests passed"
    else
        warning "Some tests failed, but continuing..."
    fi
}

# 构建项目
build_project() {
    info "Building project..."
    make build-examples
    success "Build completed"
}

# 显示帮助
show_help() {
    echo "AData-Go Quick Start Script"
    echo ""
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  setup     - Setup development environment"
    echo "  test      - Run tests"
    echo "  build     - Build all examples"
    echo "  run       - Run basic example"
    echo "  clean     - Clean build artifacts"
    echo "  docker    - Build and run Docker container"
    echo "  help      - Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 setup     # Setup environment and run tests"
    echo "  $0 run       # Run the basic example"
    echo "  $0 docker    # Build and run with Docker"
}

# 设置开发环境
setup_environment() {
    info "Setting up development environment..."
    check_go
    check_dependencies
    
    # 安装开发工具
    info "Installing development tools..."
    if command -v golangci-lint &> /dev/null; then
        info "golangci-lint already installed"
    else
        info "Installing golangci-lint..."
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    fi
    
    run_tests
    build_project
    success "Development environment setup completed"
}

# 运行基础示例
run_basic_example() {
    info "Running basic example..."
    if [ -f "bin/examples/basic" ]; then
        ./bin/examples/basic
    else
        info "Binary not found, building first..."
        make build-examples
        ./bin/examples/basic
    fi
}

# 清理
clean_project() {
    info "Cleaning project..."
    make clean
    success "Clean completed"
}

# Docker构建和运行
docker_run() {
    info "Building Docker image..."
    docker build -t adata-go:latest .
    
    info "Running Docker container..."
    docker run --rm -it adata-go:latest
}

# 主函数
main() {
    case "${1:-setup}" in
        setup)
            setup_environment
            ;;
        test)
            check_go
            check_dependencies
            run_tests
            ;;
        build)
            check_go
            check_dependencies
            build_project
            ;;
        run)
            check_go
            run_basic_example
            ;;
        clean)
            clean_project
            ;;
        docker)
            docker_run
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            error "Unknown command: $1"
            show_help
            exit 1
            ;;
    esac
}

# 检查是否在正确的目录
if [ ! -f "go.mod" ]; then
    error "Please run this script from the project root directory"
    exit 1
fi

# 运行主函数
main "$@"
@echo off
REM AData-Go Windows 启动脚本

setlocal EnableDelayedExpansion

REM 颜色定义（简化版）
set "INFO=[INFO]"
set "SUCCESS=[SUCCESS]"
set "WARNING=[WARNING]"
set "ERROR=[ERROR]"

:main
if "%1"=="" (
    call :setup_environment
) else if "%1"=="setup" (
    call :setup_environment
) else if "%1"=="test" (
    call :run_tests
) else if "%1"=="build" (
    call :build_project
) else if "%1"=="run" (
    call :run_basic_example
) else if "%1"=="clean" (
    call :clean_project
) else if "%1"=="help" (
    call :show_help
) else (
    echo %ERROR% Unknown command: %1
    call :show_help
    exit /b 1
)
goto :eof

:check_go
echo %INFO% Checking Go installation...
go version >nul 2>&1
if errorlevel 1 (
    echo %ERROR% Go is not installed. Please install Go 1.18 or later.
    exit /b 1
)
for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo %INFO% Go version: %GO_VERSION%
goto :eof

:check_dependencies
echo %INFO% Checking dependencies...
go mod download
go mod tidy
echo %SUCCESS% Dependencies checked
goto :eof

:run_tests
call :check_go
call :check_dependencies
echo %INFO% Running tests...
go test -v ./...
if errorlevel 1 (
    echo %WARNING% Some tests failed, but continuing...
) else (
    echo %SUCCESS% All tests passed
)
goto :eof

:build_project
call :check_go
call :check_dependencies
echo %INFO% Building project...

REM 创建bin目录
if not exist "bin\examples" mkdir "bin\examples"

REM 构建示例
echo %INFO% Building basic example...
go build -o bin\examples\basic.exe .\examples\basic

echo %INFO% Building concurrent example...
go build -o bin\examples\concurrent.exe .\examples\concurrent

echo %INFO% Building stock_info example...
go build -o bin\examples\stock_info.exe .\examples\stock_info

echo %SUCCESS% Build completed
goto :eof

:setup_environment
echo %INFO% Setting up development environment...
call :check_go
call :check_dependencies
call :run_tests
call :build_project
echo %SUCCESS% Development environment setup completed
goto :eof

:run_basic_example
echo %INFO% Running basic example...
if exist "bin\examples\basic.exe" (
    bin\examples\basic.exe
) else (
    echo %INFO% Binary not found, building first...
    call :build_project
    bin\examples\basic.exe
)
goto :eof

:clean_project
echo %INFO% Cleaning project...
if exist "bin" rmdir /s /q "bin"
if exist "coverage.out" del "coverage.out"
if exist "coverage.html" del "coverage.html"
go clean -cache
echo %SUCCESS% Clean completed
goto :eof

:show_help
echo AData-Go Quick Start Script (Windows)
echo.
echo Usage: %~nx0 [command]
echo.
echo Commands:
echo   setup     - Setup development environment
echo   test      - Run tests
echo   build     - Build all examples
echo   run       - Run basic example
echo   clean     - Clean build artifacts
echo   help      - Show this help message
echo.
echo Examples:
echo   %~nx0 setup     # Setup environment and run tests
echo   %~nx0 run       # Run the basic example
goto :eof

REM 检查是否在正确的目录
if not exist "go.mod" (
    echo %ERROR% Please run this script from the project root directory
    exit /b 1
)
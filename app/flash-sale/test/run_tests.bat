@echo off
echo Starting Flash Sale System Tests...
echo.

REM 设置Go环境
set GO111MODULE=on
set GOPROXY=https://goproxy.io,direct

REM 进入测试目录
cd /d c:\code\shop-goframe-micro-service-refacotor\app\flash-sale

echo 1. Running unit tests...
go test -v ./test/... -run TestRateLimiter
echo.

echo 2. Running anti-brush tests...
go test -v ./test/... -run TestAntiBrushChecker
echo.

echo 3. Running integration tests...
go test -v ./test/... -run TestFlashSaleSystem
echo.

echo 4. Running performance benchmark...
go test -bench=. -benchmem ./test/...
echo.

echo 5. Running code coverage analysis...
go test -coverprofile=coverage.out ./test/...
go tool cover -html=coverage.out -o coverage.html
echo Coverage report generated: coverage.html
echo.

echo All tests completed!
echo.
echo Test Summary:
echo - Unit tests: Verify individual components
echo - Integration tests: Verify system workflow
echo - Performance tests: Benchmark system performance
echo - Coverage analysis: Code coverage report
echo.
echo Check test results above for any failures.
pause
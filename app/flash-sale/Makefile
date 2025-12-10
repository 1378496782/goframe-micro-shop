PACKAGE = c:\\code\\shop-goframe-micro-service-refacotor\\app\\flash-sale

VERSION = v1.0.0

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building binary..."
	@go build -o ./bin/flash-sale main.go

.PHONY: run
run:
	@echo "Running application..."
	@go run main.go

.PHONY: clean
clean:
	@echo "Cleaning binary..."
	@rm -f ./bin/flash-sale

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: docker
docker:
	@echo "Building docker image..."
	@docker build -t flash-sale:$(VERSION) -f manifest/docker/Dockerfile .
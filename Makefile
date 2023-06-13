DOCKER ?= docker
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

test: export CORDLE_CONFIG_PATH=test_config.json

# clean & dev

clean:
	@rm -r build

dev:
	@bash scripts/dev.sh

# testing

test:
	$(GO) clean -testcache 
	$(GO) mod tidy
	$(GO) test -cover -v ./...

# deploy & build

build: 
	$(GO) build -o build/program/app cmd/cli/main.go 

tdeploy:
	$(DOCKER) build --tag cordle2 .
	$(DOCKER) run -rm cordle2

deploy:
	$(DOCKER) build --tag cordle2 .
	$(DOCKER) run -d cordle2

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: dev clean test build tdeploy deploy fmt mysql
DOCKER ?= docker
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

# clean & dev

clean:
	@rm -r build

dev:
	@bash scripts/dev.sh

# testing

test:
	$(GO) clean -testcache 
	$(GO) mod tidy
	CORDLE_CONFIG_PATH=test_config.json $(GO) test -cover ./... -tags=unit

testint:
	$(DOCKER) compose -f docker-compose.test.yml build
	$(DOCKER) compose -f docker-compose.test.yml up --abort-on-container-exit
	$(DOCKER) compose down

# deploy & build

build: 
	$(GO) build -o build/program/app cmd/cli/main.go 

tdeploy:
	$(DOCKER) build --tag cordle2 .
	$(DOCKER) run -rm cordle2

deploy:
	$(DOCKER) compose build
	$(DOCKER) compose up -d

shutdown:
	$(DOCKER) compose down

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

mysql:
	cd deployment/mysql; make deploy

.PHONY: dev clean test build tdeploy deploy fmt mysql
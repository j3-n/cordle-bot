## general

.PHONY: dev clean

clean:
	@rm -r build

dev:
	@bash scripts/dev.sh

## testing

.PHONY: test

test:
	go clean -testcache 
	go mod tidy
	go test -cover ./...

## deploy

.PHONY: build tdeploy deplo

build: 
	go build -o build/program/app cmd/cli/main.go 

tdeploy:
	docker build --tag cordle2 .
	docker run -rm cordle2 &

deploy:
	docker build --tag cordle2 .
	docker run cordle2 &
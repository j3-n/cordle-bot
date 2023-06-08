DOCKER ?= docker

deploy:
	$(DOCKER) build --tag mysql .
	$(DOCKER) run -it -p 3306:3306 mysql

.PHONY: deploy
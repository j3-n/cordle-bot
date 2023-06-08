DOCKER ?= docker

deploy:
    $(DOCKER) build -t mysql_db .
    $(DOCKER) run mysql_db

.PHONY: deploy
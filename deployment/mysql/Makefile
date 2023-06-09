DOCKER ?= docker

deploy:
	$(DOCKER) build --tag mysql_db .
	$(DOCKER) run mysql_db

.PHONY: deploy
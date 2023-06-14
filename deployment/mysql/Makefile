DOCKER ?= docker

deploy:
	$(DOCKER) build --tag mysql_db .
	$(DOCKER) run -d --name mysql_app mysql_db

.PHONY: deploy
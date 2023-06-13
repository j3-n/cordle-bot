DOCKER ?= docker

deploy:
	$(DOCKER) build --tag mysql_db .
	$(DOCKER) run -p 3306:9876 -d --name mysql_app mysql_db

.PHONY: deploy
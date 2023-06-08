deploy:
	$(DOCKER) build --tag mysql -f deployment/mysql/Dockerfile .
	$(DOCKER) run -it -p 3306:3306 mysql

.PHONY: deploy
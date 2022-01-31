# include help.mk

.PHONY: version clean lint install build image tag push release run run-local  remove-docker
.DEFAULT_GOAL := help

GITHUB_GROUP = esequielvirtuoso
HUB_HOST     = hub.docker.com
HUB_USER 	 = esequielvirtuoso
HUB_GROUP    =

BUILD         	= $(shell git rev-parse --short HEAD)
DATE          	= $(shell date -uIseconds)
VERSION  	  	= $(shell git describe --always --tags)
NAME           	= $(shell basename $(CURDIR))
IMAGE          	= $(HUB_HOST)/$(HUB_USER)/$(HUB_GROUP)/$(NAME):$(BUILD)

MYSQL_NAME = mysql_$(NAME)_$(BUILD)
MYSQL_ADMINER_NAME = mysql_adminer_$(NAME)_$(BUILD)
NETWORK_NAME  = network_$(NAME)_$(BUILD)

git-config:
	git config --replace-all core.hooksPath .githooks

check-env-%:
	@ if [ "${${*}}" = ""  ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

version: ##@other Check version.
	@echo $(VERSION)

clean: ##@dev Remove folder vendor, public and coverage.
	rm -rf vendor public coverage

install: clean ##@dev Download dependencies via go mod.
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

audit: ##@check Run vulnerability check in Go dependencies.
	DOCKER_BUILDKIT=1 docker build --progress=plain --target=audit --file=Dockerfile .

lint: ##@check Run lint on docker.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--target=lint \
		--file=Dockerfile .

env: ##@environment Create network and run mysql container.
	MYSQL_NAME=${MYSQL_NAME} \
	MYSQL_ADMINER_NAME=$(MYSQL_ADMINER_NAME) \
	NETWORK_NAME=${NETWORK_NAME} \
	docker-compose up -d

env-ip: ##@environment Return local MYSQL IP (from Docker container)
	@echo $$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${MYSQL_NAME})

env-stop: ##@environment Remove mysql container and remove network.
	MYSQL_NAME=${MYSQL_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose kill
	MYSQL_NAME=${MYSQL_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose rm -vf
	MYSQL_ADMINER_NAME=${MYSQL_ADMINER_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose kill
	MYSQL_ADMINER_NAME=${MYSQL_ADMINER_NAME} NETWORK_NAME=${NETWORK_NAME} docker-compose rm -vf
	docker network rm $(NETWORK_NAME)

run:
	go run main.go

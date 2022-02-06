# include help.mk

.PHONY: version clean lint install build image tag push release run run-local  remove-docker
.DEFAULT_GOAL := help

GITHUB_GROUP = esequielvirtuoso
HUB_HOST     = hub.docker.com/repository/docker
HUB_USER 	 = esequielvirtuoso
HUB_GROUP    = go_apps

BUILD         	= $(shell git rev-parse --short HEAD)
DATE          	= $(shell date -uIseconds)
VERSION  	  	= $(shell git describe --always --tags)
NAME           	= $(shell basename $(CURDIR))
IMAGE          	= $(HUB_HOST)/$(HUB_USER)/$(HUB_GROUP)/$(NAME):$(BUILD)

MYSQL_NAME = mysql_$(NAME)_$(BUILD)
MYSQL_ADMINER_NAME = mysql_adminer_$(NAME)_$(BUILD)
NETWORK_NAME  = network_$(NAME)_$(BUILD)
MYSQL_URL = root:passwd@tcp(127.0.0.1:3305)/users_db?charset=utf8

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

test: ##@check Run tests and coverage.
	docker build --progress=plain \
		--network $(NETWORK_NAME) \
		--tag $(IMAGE) \
		--build-arg MYSQL_URL=$(MYSQL_URL) \
		--target=test \
		--file=Dockerfile .

	-mkdir coverage
	docker create --name $(NAME)-$(BUILD) $(IMAGE)
	docker cp $(NAME)-$(BUILD):/index.html ./coverage/.
	docker rm -vf $(NAME)-$(BUILD)

build: ##@build Build image.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(IMAGE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD=$(BUILD) \
		--build-arg DATE=$(DATE) \
		--target=build \
		--file=Dockerfile .

image: check-env-VERSION ##@build Create release docker image.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--tag $(IMAGE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD=$(BUILD) \
		--build-arg DATE=$(DATE) \
		--target=image \
		--file=Dockerfile .

tag: check-env-VERSION ##@build Add docker tag.
	docker tag $(IMAGE) \
		$(HUB_HOST)/$(HUB_USER)/$(HUB_GROUP)/$(NAME):$(VERSION)

push: check-env-VERSION ##@build Push docker image to registry.
	docker push $(HUB_HOST)/$(HUB_USER)/$(HUB_GROUP)/$(NAME):$(VERSION)

release: check-env-TAG ##@build Create and push git tag.
	git tag -a $(TAG) -m "Generated release "$(TAG)
	git push origin $(TAG)

run:
	go run main.go

run-local: ##@dev Run locally.
	LOGGER_LEVEL=debug \
	MYSQL_URL=root:passwd@$$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(MYSQL_NAME):3305)/users_db?charset=utf8 \
	run

run-docker: check-env-POSTGRES_URL ##@docker Run docker container.
	docker run --rm \
		--name $(NAME) \
		-e LOGGER_LEVEL=debug \
		-e MYSQL_URL=$(MYSQL_URL) \
		-p 5001:8080 \
		$(IMAGE)

remove-docker: ##@docker Remove docker container.
	-docker rm -vf $(NAME)

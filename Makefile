DOCKER := docker-compose -f deployments/docker-compose.yml
PROTOC := ${DOCKER} run --rm go-protoc protoc
GO := ${DOCKER} run go-protoc go
CMDS := calc_client calc_server

export PROTOC
export GO
export CMDS

.PHONY: all
all: test build

.PHONY: test
test:
	scripts/test.sh

.PHONY: build
build:
	scripts/build.sh

.PHONY: push
push:
	scripts/push.sh

.PHONY: start
start:
	${DOCKER} up

#!/usr/bin/env bash

CURDIR=$(pwd)
TAG=shadmanx/calc-server:latest

docker build --pull --target server ${CURDIR}/build/package -t ${TAG}
docker push ${TAG}

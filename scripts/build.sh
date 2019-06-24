#!/usr/bin/env bash

SCRIPTSDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
${SCRIPTSDIR}/protoc.sh

# Build all binaries
for CMD in ${CMDS}; do
    OUT="build/package/${CMD}"
    ${GO} build -a -tags netgo -o ${OUT} cmd/${CMD}/main.go
done

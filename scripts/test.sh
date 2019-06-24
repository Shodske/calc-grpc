#!/usr/bin/env bash

SCRIPTSDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
${SCRIPTSDIR}/protoc.sh

${GO} test ./...


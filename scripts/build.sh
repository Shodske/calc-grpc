#!/usr/bin/env bash

CURDIR=$(pwd)

# Generate .pb.go files from all .proto files
for FILE in $(find ${CURDIR} -name '*.proto'); do
    [[ -f "$FILE" ]] || break

    FILEPATH=$(realpath --relative-to="$CURDIR" "$FILE")
    PKG=$(dirname "$FILEPATH")

    ${PROTOC} -I ${PKG}/ ${FILEPATH} --go_out=plugins=grpc:${PKG}
done

# Build all binaries
for CMD in ${CMDS}; do
    OUT="build/package/${CMD}"
    ${GO} build -a -tags netgo -o ${OUT} cmd/${CMD}/main.go
done

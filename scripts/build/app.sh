#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"
source "${SCRIPTS_DIR}/helpers-source.sh"

BIN_DIR=${REPO_ROOT}/bin
mkdir -p "${BIN_DIR}"

echo "${SCRIPT_NAME} is running... "

APP=${APP_NAME}

echo "Building ${APP}..."

BIN_OUT="${BIN_DIR}/${APP}"

GO_BUILD_PACKAGE="${REPO_ROOT}/cmd/${APP}"

rm -rf "${BIN_OUT}"

go build -trimpath -o "${BIN_OUT}" "${GO_BUILD_PACKAGE}"

echo "Build ${BIN_OUT} success"

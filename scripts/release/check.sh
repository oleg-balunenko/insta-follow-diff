#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

APP=${APP_NAME}

echo "${SCRIPT_NAME} is running fo ${APP}... "

checkInstalled 'goreleaser'

goreleaser healthcheck

goreleaser check

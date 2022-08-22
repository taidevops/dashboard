#!/usr/bin/env bash

# Directories.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"

TMP_DIR="${ROOT_DIR}/.tmp"
FRONTEND_DIR="${TMP_DIR}/frontend"

# Binaries.
NG_BIN="${ROOT_DIR}/node_modules/.bin/ng"

function say { echo "$@"; }

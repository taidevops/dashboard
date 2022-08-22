#!/usr/bin/env bash

# Exit on error
set -e

# Import config.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"
. "${ROOT_DIR}/aio/scripts/conf.sh"

# Declare variables.
FRONTEND_ONLY=false

function clean {
  rm -rf ${DIST_DIR} ${TMP_DIR}
}

function build::frontend {
  say "\nBuilding localized frontend"
  mkdir -p ${FRONTEND_DIR}
  ${NG_BIN} build \
            --configuration production
}

if [ "${FRONTEND_ONLY}" = true ] ; then
  build::frontend
  exit
fi

build::frontend

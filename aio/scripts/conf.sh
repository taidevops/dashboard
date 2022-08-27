#!/usr/bin/env bash

# Directories.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"

TMP_DIR="${ROOT_DIR}/.tmp"
FRONTEND_DIR="${TMP_DIR}/frontend"

CACHE_DIR="${ROOT_DIR}/.cached_tools"

# Binaries.
NG_BIN="${ROOT_DIR}/node_modules/.bin/ng"

# Global constants.
ARCH=$(uname | awk '{print tolower($0)}')

# Local cluster configuration (check start-cluster.sh script for more details).
KIND_VERSION="v0.14.0"
K8S_VERSION="v1.24.3"
KIND_BIN=${CACHE_DIR}/kind-${KIND_VERSION}
CODEGEN_VERSION="v0.24.1"
CODEGEN_BIN=${GOPATH}/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}/generate-groups.sh

function say { echo "$@"; }

function ensure-cache {
  say "\nMaking sure that ${CACHE_DIR} directory exists"
  mkdir -p ${CACHE_DIR}
}

function download-kind {
  KIND_URL="https://github.com/kubernetes-sigs/kind/releases/download/${KIND_VERSION}/kind-${ARCH}-amd64"
  say "\nDownloading kind ${KIND_URL} if it is not cached"
  wget -nc -O ${KIND_BIN} ${KIND_URL}
  chmod +x ${KIND_BIN}
  ${KIND_BIN} version
}

function ensure-kubeconfig {
  say "\nMaking sure that kubeconfig file exists and will be used by Dashboard"
  mkdir -p ${HOME}/.kube
  touch ${HOME}/.kube/config

  # Let's back up the kubeconfig so we don't totally blow it away
  # I learned from personal experience. It made me sad. :(
  # ${HOME}/.kube/config is mounted in container for development,
  # so we can not `mv` or `rm` it.
  cp ${HOME}/.kube/config ${HOME}/.kube/config-unkind

  ${KIND_BIN} get kubeconfig --name="k8s-cluster-ci" > ${HOME}/.kube/config
}

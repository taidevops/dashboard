# Import config.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"
. "${ROOT_DIR}/aio/scripts/conf.sh"

function start-ci-heapster {
  say "\nRunning heapster in standalone mode"

  say "\nWaiting for heapster to be started"
}

function start-kind {
  ${KIND_BIN} create cluster --name="k8s-cluster-ci" --image="kindest/node:${K8S_VERSION}"
  ensure-kubeconfig

  say "\nKubernetes cluster is ready to use"
}

# Execute script.
ensure-cache
download-kind
start-kind

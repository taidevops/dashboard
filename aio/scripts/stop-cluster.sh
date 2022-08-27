# Import config.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"
. "${ROOT_DIR}/aio/scripts/conf.sh"

ensure-cache
download-kind

${KIND_BIN} delete cluster --name="k8s-cluster-ci"

# Restore the original kubeconfig and all's right
# with the world.
# ${HOME}/.kube/config is mounted in container for development,
# so we can not `mv` or `rm` it.
cat ${HOME}/.kube/config-unkind > ${HOME}/.kube/config
rm -f ${HOME}/.kube/config-unkind
set -o errexit
set -o nounset
set -o pipefail

# Import config.
ROOT_DIR="$(cd $(dirname "${BASH_SOURCE}")/../.. && pwd -P)"
source "${ROOT_DIR}/aio/scripts/conf.sh"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
"${CODEGEN_BIN}" "deepcopy,client,informer,lister" \
  github.com/taidevops/dashboard/src/app/backend/plugin/client github.com/taidevops/dashboard/src/app/backend/plugin \
  apis:v1alpha1 \
  --go-header-file "${ROOT_DIR}"/aio/scripts/license-header.go.txt \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/.."

# Remove old generated client
rm -rf ./src/app/backend/plugin/client
# Move generated deepcopy funcs and client
mv "$(dirname "${BASH_SOURCE[0]}")"/../github.com/taidevops/dashboard/src/app/backend/plugin/apis/v1alpha1/zz_generated.deepcopy.go ./src/app/backend/plugin/apis/v1alpha1
mv "$(dirname "${BASH_SOURCE[0]}")"/../github.com/taidevops/dashboard/src/app/backend/plugin/client ./src/app/backend/plugin
# Remove empty directory
rm -rf ./aio/github.com

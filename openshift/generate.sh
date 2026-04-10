#!/usr/bin/env bash

set -euo pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

echo "Clean up unused upstream actions"
rm -f "${repo_root_dir}/.github/workflows/kind-e2e-upgrade.yaml"

go run github.com/openshift-knative/hack/cmd/generate@latest \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --includes cmd \
  --app-file-fmt "/ko-app/%s" \
  --dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.25-openshift-4.21"

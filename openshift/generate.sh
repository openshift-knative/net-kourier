#!/usr/bin/env bash

set -euo pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

go run github.com/openshift-knative/hack/cmd/generate@latest \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --includes cmd \
  --app-file-fmt "/ko-app/%s" \
  --dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-8-release-golang-%s-openshift-4.17"

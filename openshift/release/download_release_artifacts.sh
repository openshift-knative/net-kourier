#!/usr/bin/env bash

# Usage example: ./download_release_artifacts.sh v1.8.0

set -Eeuo pipefail

SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"

# Returns the major and minor part of the whole version, joined with a dot.
function versions.major_minor {
  local version=${1:?Pass a full version as arg[1]}
  # shellcheck disable=SC2001
  # Ref: https://regex101.com/r/Po1HA3/1
  echo "${version}" | sed 's/^v\?\([[:digit:]]\+\)\.\([[:digit:]]\+\).*/\1.\2/'
}

version=$1
patches_path="${SCRIPT_DIR}/../patches"
artifacts_path="${SCRIPT_DIR}/artifacts"
mkdir -p "${patches_path}"
mkdir -p "${artifacts_path}"

url="https://github.com/knative-sandbox/net-kourier/releases/download/knative-${version}/kourier.yaml"
kourier_file="${artifacts_path}/0-kourier.yaml"
wget --no-check-certificate "$url" -O "$kourier_file"
# TODO: [SRVKS-610] These values should be replaced by operator instead of sed.
sed -i -e 's/net-kourier-controller.knative-serving/net-kourier-controller/g' "$kourier_file"

# Download config-network.yaml for Kourier. This is necessary as kourier uses different namespace (knative-serving-ingress).
config_network_url="https://raw.githubusercontent.com/knative/networking/release-$(versions.major_minor "${version}")/config/config-network.yaml"
config_network="${artifacts_path}/1-config-network.yaml"
wget --no-check-certificate "$config_network_url" -O "$config_network"
sed -i -e '/labels:$/a \    app.kubernetes.io\/component: kourier' "$config_network"
sed -i -e '/labels:$/a \    networking.knative.dev\/ingress-provider: kourier' "$config_network"

# Make Kourier rollout in a more defensive way so no requests get dropped.
# TODO: Can probably be removed in 1.21 and/or be sent upstream.
git apply "${patches_path}/001-kourier-rollout.patch"
git apply "${patches_path}/002-backport.patch"
git apply "${patches_path}/003-keepalive.patch"

#!/usr/bin/env bash

# Usage example: ./download_release_artifacts.sh 1.8.0

set -Eeuo pipefail

SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"

# TODO: automatically detects the version via branch name or something.
VERSION=$1

function resolve_resources(){
  local dir=$1
  local resolved_file_name=$2

  echo "Writing resolved yaml to $resolved_file_name"

  > "$resolved_file_name"

  for yaml in `find $dir -name "*.yaml" | sort`; do
    resolve_file "$yaml" "$resolved_file_name"
  done
}

function resolve_file() {
  local file=$1
  local to=$2

  echo "---" >> "$to"

  echo $file

  sed -e "s+app.kubernetes.io/version: devel+app.kubernetes.io/version: \""$VERSION"\"+" \
      "$file" >> "$to"

}

readonly YAML_OUTPUT_DIR="openshift/release/artifacts/"
readonly KOURIER_YAML=${YAML_OUTPUT_DIR}/0-kourier.yaml
readonly CONFIG_NETWORK_YAML=${YAML_OUTPUT_DIR}/1-config-network.yaml

# Clean up
rm -rf "$YAML_OUTPUT_DIR"
mkdir -p "$YAML_OUTPUT_DIR"

readonly CONFIG_NETWORK="vendor/knative.dev/networking/config/config-network.yaml"

patches_path="${SCRIPT_DIR}/../patches"

# TODO: [SRVKS-610] These values should be replaced by operator instead of sed.
git apply "${patches_path}/001-service-location.patch"

# Add label to the config-network, which is deployed in knative-serving-ingress namespace.
sed -i -e 's/app.kubernetes.io\/component: networking/app.kubernetes.io\/component: kourier\n    networking.knative.dev\/ingress-provider: kourier/' "$CONFIG_NETWORK"

resolve_resources "config/" "$KOURIER_YAML"
resolve_resources "vendor/knative.dev/networking/config/config-network.yaml" "$CONFIG_NETWORK_YAML"

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
readonly KOURIER_YAML=${YAML_OUTPUT_DIR}/net-kourier.yaml
readonly patches_path="${SCRIPT_DIR}/../patches"

# Clean up
rm -rf "$YAML_OUTPUT_DIR"
mkdir -p "$YAML_OUTPUT_DIR"
# clean up before applying patch
git apply -R "${patches_path}"/* || true

git apply "${patches_path}"/*

resolve_resources "openshift/release/extra/ config" "$KOURIER_YAML"

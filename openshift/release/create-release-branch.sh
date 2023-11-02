#!/usr/bin/env bash

# Usage: create-release-branch.sh release-1.9
#
#
# The repository assumes that you have the following branch names.
#
#   $ git remote -v
#   upstream	git@github.com:knative-sandbox/net-kourier.git (fetch)
#   upstream	git@github.com:knative-sandbox/net-kourier.git (push)
#   openshift	git@github.com:openshift-knative/net-kourier.git (fetch)
#   openshift	git@github.com:openshift-knative/net-kourier.git (push)
#
set -e # Exit immediately on error.

release=$1

# Set upstream release without "v" prefix. e.g. release-v1.11 => release-1.11
upstream_release=release-"${release#"release-v"}"

# Fetch the latest upstream and checkout the new branch.
git fetch upstream "${upstream_release}"
git checkout upstream/"${upstream_release}"

# Copy the openshift extra files from the OPENSHIFT/main branch.
git fetch openshift main
git checkout openshift/main -- openshift OWNERS
git add openshift OWNERS
git commit -m "Add openshift specific files."

openshift/release/download_release_artifacts.sh "${release}"
git add .
git commit -am ":fire: Generate artifacts."

# TODO: currently this script is executed manually. So, do not push automatically.
echo "
Now ready to create a new branch. Push it by:

  $ git checkout -b ${release}
  $ git push openshift ${release}

"

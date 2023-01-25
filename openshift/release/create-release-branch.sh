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

# Fetch the latest upstream and checkout the new branch.
git fetch upstream ${release}
git checkout ${release}

# Copy the openshift extra files from the OPENSHIFT/main branch.
git fetch openshift main
git checkout openshift/main -- openshift OWNERS Dockerfile
git add openshift OWNERS Dockerfile
git commit -m "Add openshift specific files."

# Drop the release-suffix and add the micro version with '0'.
# e.g. release-1.9 => 1.9.0
VERSION=${release#"release-"}.0

openshift/release/download_release_artifacts.sh ${VERSION}
git add .
git commit -am ":fire: Generate artifacts."

# TODO: currently this script is executed manually. So, do not push automatically.
echo "
Now ready to create a new branch. Push it by:

  $ git push openshift ${release}

"

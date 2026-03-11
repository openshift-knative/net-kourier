#!/usr/bin/env bash

set -euo pipefail

TARGET_FILE=$1

# PDB minAvailable: set '80%' to '1'
sed -i 's/minAvailable: 80%/minAvailable: 1/' "$TARGET_FILE"

# Drop explicit runAsUser/runAsGroup (let OpenShift SCC assign them)
sed -i '/^\s*runAsUser: 65534$/d' "$TARGET_FILE"
sed -i '/^\s*runAsGroup: 65534$/d' "$TARGET_FILE"

# SRVKS-1343 set soft ulimit -n to the value of the hard limit for envoy
sed -i '/command:/{N;s|command:\n\(\s*\)- /usr/local/bin/envoy|command:\n\1["/bin/sh", "-c", "ulimit -S -n $(ulimit -n -H); exec /usr/local/bin/envoy \\"$@\\""]|}' "$TARGET_FILE"


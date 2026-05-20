generate-release:
# Run both update-codegen.sh and update-deps.sh to reconcile deps on update.
# There are no vendor/ changing patches in this repo, otherwise we'd need to re-apply those.
	./hack/update-codegen.sh
	./openshift/generate.sh
.PHONY: generate-release

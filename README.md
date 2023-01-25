# OpenShift net-kourier

This repository holds OpenShift's fork of
[`knative/net-kourier`](https://github.com/knative-sandbox/net-kourier) with additions and
fixes needed only for the OpenShift side of things.

# OpenShift net-kourier Release procedure

Currently the release cut is not automated but just one command.
Once upstream cuts the branch (e.g. upstream cut `release-1.9` branch), run:

```sh
$ openshift/release/create-release-branch.sh release-1.9
```

Then, push the branch against to this repository.

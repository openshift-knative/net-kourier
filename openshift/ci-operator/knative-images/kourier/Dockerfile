# DO NOT EDIT! Generated Dockerfile for cmd/kourier.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kourier

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

COPY --from=builder /usr/bin/main /ko-app/kourier
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-net-kourier-kourier-rhel8-container" \
      name="openshift-serverless-1/net-kourier-kourier-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Net Kourier Kourier" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Net Kourier Kourier" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Net Kourier Kourier" \
      io.k8s.description="Red Hat OpenShift Serverless Net Kourier Kourier" \
      io.openshift.tags="kourier"

ENTRYPOINT ["/ko-app/kourier"]

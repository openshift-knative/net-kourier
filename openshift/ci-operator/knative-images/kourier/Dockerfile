FROM registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17 AS builder
WORKDIR /app/
COPY . .
RUN go build -mod vendor -o /tmp/kourier ./cmd/kourier

FROM openshift/origin-base
USER 65532

COPY --from=builder /tmp/kourier /ko-app/kourier
ENTRYPOINT ["/ko-app/kourier"]

FROM openshift/origin-base

USER 1001

ADD kourier /app/kourier
EXPOSE 18000 19000
ENTRYPOINT ["/app/kourier"]

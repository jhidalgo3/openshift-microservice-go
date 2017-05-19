FROM alpine:3.5
COPY openshift-microservice-go /openshift-microservice-go
EXPOSE 8080
WORKDIR /
CMD ["/openshift-microservice-go"]

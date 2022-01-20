FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/upm/udevs_go_auth_service 
WORKDIR $GOPATH/src/gitlab.udevs.io/upm/udevs_go_auth_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/udevs_go_auth_service /

FROM alpine
COPY --from=builder udevs_go_auth_service .
ENTRYPOINT ["/udevs_go_auth_service"]

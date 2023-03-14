FROM golang:1.19.5-bullseye as builder
RUN mkdir -p /go/src/github.com/EcoG-io/mender-artifact
WORKDIR /go/src/github.com/EcoG-io/mender-artifact
ADD ./ .
RUN make get-build-deps && \
    make build && \
    make install
ENTRYPOINT [ "/go/bin/mender-artifact" ]

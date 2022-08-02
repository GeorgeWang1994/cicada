FROM golang:1.17-alpine as build
USER root

ARG PRO
ENV PROJ_PATH=${GOPATH}/src/cicada/module/${PRO} \
    PKG_PATH=${GOPATH}/src/cicada/module/pkg \
    PROTO_PATH=${GOPATH}/src/cicada/module/proto
COPY ./module/${PRO} ${PROJ_PATH}
COPY ./module/pkg ${PKG_PATH}
COPY ./module/proto ${PROTO_PATH}

ENV VERSION=0.0.1

WORKDIR ${GOPATH}/src/cicada/module/${PRO}

RUN apk add --no-cache ca-certificates bash git && \
    go mod tidy && GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$(VERSION)" \
    		-o /tmp/cicada-${PRO} ./cmd

FROM alpine:latest

ARG PRO
COPY --from=build /tmp/cicada-${PRO} /app/cicada-${PRO}
COPY ./module/${PRO}/config.json /app/config.json

WORKDIR /app/

CMD ["cicada-${PRO}", "daemon", "-c", "/app/config.json"]

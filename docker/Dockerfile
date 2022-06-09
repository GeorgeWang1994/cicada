FROM golang:1.17-alpine
USER root

ARG PRO
ENV PROJ_PATH=${GOPATH}/src/cicada/${PRO}
COPY ./${PRO} ${PROJ_PATH}
ENV VERSION=0.0.1

WORKDIR ${GOPATH}/src/cicada/${PRO}

RUN apk add --no-cache ca-certificates bash git && \
    GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$(VERSION)" \
    		-o /tmp/cicada-${PRO} ./cmd

CMD ["cicada-${PRO}", "daemon"]

FROM golang:1.17-alpine
USER root

ARG PRO
ENV PROJ_PATH=${GOPATH}/src/cicada/${PRO}

RUN mkdir -p /tmp && \
    mkdir -p /tmp/logs && \
    apk add --no-cache ca-certificates bash git supervisor

COPY . ${PROJ_PATH}
RUN make all \
    && make agentd \
    && rm -rf ${PROJ_PATH}

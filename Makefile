#!/bin/bash
CMD=( agentd alarm ev judge portal )
VERSION := "0.0.1"
Module = "alarm" "ev" "judge" "portal"

.PHONY: $(CMD)
 $(CMD):
	mkdir -p ./bin/$@ ;
	cd ./module/$@ && go mod tidy && GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$VERSION" \
		-o ../../bin/$@/cicada-$@ ./cmd && cp config.json ../../bin/$@/ && cd -;

.PHONY : proto
proto:
	@echo generate proto...;
	@ for i in module/proto/*/*.proto; \
	do \
		/usr/local/bin/protoc --proto_path=. --go_out=proto/ --go-grpc_out=proto/ $$i; \
	done

.PHONY: all
all:
	@ for i in $(Module); do \
    	make $$i; \
    done

.PHONY: build_all
build_all:
	@ for i in $(Module); do \
	  	make $$i ; \
	  	docker build --build-arg PRO=$$i -t cicada-$$i:latest -f docker/Dockerfile . ; \
	done

clean:
	@echo clean bin...
	@rm -rf ./bin

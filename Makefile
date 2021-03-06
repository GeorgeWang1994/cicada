#!/bin/bash
CMD=( agentd alarm ev judge portal )
VERSION := $(shell cat VERSION)

.PHONY: $(CMD)
 $(CMD):
	mkdir -p ./bin/$@ ;
	go mod tidy;
	cd ./$@ && GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$(VERSION)" \
		-o ../bin/$@/cicada-$@ ./cmd ;

.PHONY : proto
proto:
	@echo generate proto...;
	@ for i in proto/*/*.proto; \
	do \
		/usr/local/bin/protoc --proto_path=. --go_out=proto/ --go-grpc_out=proto/ $$i; \
	done

.PHONY: all
all:
	make agentd;
	make alarm;
	make ev;
	make judge;
	make portal;


.PHONY: build
build:
	@ ARR=( agentd alarm ev judge portal )
	@ echo ${ARR[@]}
	@ for i in ${ARR[@]}; \
	do \
	  	echo $i...; \
	done

clean:
	@echo clean bin...
	@rm -rf ./bin
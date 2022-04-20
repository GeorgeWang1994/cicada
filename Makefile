CMD = agentd alarm cron ev judge portal
VERSION := $(shell cat VERSION)

.PHONY: $(CMD)
 $(CMD):
	mkdir -p ./bin/$@ ;
	GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$(VERSION)" \
		-o bin/$@/cicada-$@ ./$@/cmd ;

.PHONY : proto
proto:
	@echo generate proto...;
	@ for i in proto/*/*.proto; \
	do \
		/usr/local/bin/protoc --proto_path=. --go_out=proto/ --go-grpc_out=proto/ $$i; \
	done

clean:
	@echo clean bin...
	@rm -rf ./bin

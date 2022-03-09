CMD = agentd alarm cron ev evadm
VERSION := $(shell cat VERSION)

.PHONY: $(CMD)
 $(CMD):
	mkdir -p ./bin/$@ ;
	GO111MODULE=on go build -ldflags "-X main.BinaryName=$@ -X main.GitCommit=`git rev-parse --short HEAD` -X main.Version=$(VERSION)" \
		-o bin/$@/cicada-$@ ./$@/cmd ;

.PHONY:
docker:
	@if [ -e out ] ; then rm -rf out; fi
	@mkdir out
	@$(foreach var,$(CMD),mkdir -p ./out/$(var)/bin;)
	@$(foreach var,$(CMD),mkdir -p ./out/$(var)/config;)
	@$(foreach var,$(CMD),mkdir -p ./out/$(var)/logs;)
	@$(foreach var,$(CMD),cp ./config/$(var).json ./out/$(var)/config/cfg.json;)
	@$(foreach var,$(CMD),cp ./bin/$(var)/falcon-$(var) ./out/$(var)/bin;)
	@if expr "$(CMD)" : "agent" > /dev/null; then \
		(cp -r ./modules/agent/public ./out/agent/); \
		(cd ./out && ln -s ./agent/public ./public); \
		(cd ./out && mkdir -p ./agent/plugin && ln -s ./agent/plugin ./plugin); \
	fi
	@cp ./docker/ctrl.sh ./out/ && chmod +x ./out/ctrl.sh
	@cp $(TARGET) ./out/$(TARGET)
	tar -C out -zcf open-falcon-v$(VERSION).tar.gz .
	@rm -rf out


.PHONY : proto
proto:
	@echo generate proto...;
	@ for i in proto/*.proto; \
	do \
		protoc --proto_path=. --go_out=proto/ $$i; \
	done

clean:
	@echo clean bin...
	@rm -rf ./bin

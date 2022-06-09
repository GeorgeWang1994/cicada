module github.com/GeorgeWang1994/cicada/module/agentd

go 1.17

require (
	github.com/google/uuid v1.3.0
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69
	github.com/spf13/cobra v1.3.0
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.3 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/GeorgeWang1994/cicada/module/pkg => ../pkg

replace github.com/GeorgeWang1994/cicada/module/proto => ../proto

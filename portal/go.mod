module cicada/portal

go 1.17

require (
	cicada/pkg v0.0.1
	cicada/proto v0.0.1
	github.com/garyburd/redigo v1.6.3
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v1.1.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace cicada/pkg v0.0.1 => ../pkg

replace cicada/proto v0.0.1 => ../proto

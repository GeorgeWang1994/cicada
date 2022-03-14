module cicada/judge

go 1.17

require (
	cicada/pkg v0.0.1
	cicada/proto v0.0.1
	github.com/garyburd/redigo v1.6.3
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69
	github.com/segmentio/kafka-go v0.4.30
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v1.1.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace cicada/pkg v0.0.1 => ../pkg

replace cicada/proto v0.0.1 => ../proto

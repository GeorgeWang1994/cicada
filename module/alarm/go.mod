module github.com/GeorgeWang1994/cicada/module/alarm

go 1.17

require (
	github.com/GeorgeWang1994/cicada/module/pkg v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.3
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.3.0
	github.com/toolkits/net v0.0.0-20160910085801-3f39ab6fe3ce
)

require (
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/GeorgeWang1994/cicada/module/pkg => ../pkg

replace github.com/GeorgeWang1994/cicada/module/proto => ../proto

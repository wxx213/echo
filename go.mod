module github.wxx.example

go 1.15

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/kitex latest
	github.com/cloudwego/netpoll latest
	github.com/kitex-contrib/monitor-prometheus v0.0.0-20210817080809-024dd7bd51e1
	github.com/kitex-contrib/registry-zookeeper v0.0.0-20211217154151-5e91ee291af8
	google.golang.org/protobuf v1.26.0
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

// use customed packages
replace (
       github.com/cloudwego/kitex latest => ../kitex
       github.com/cloudwego/netpoll latest => ../netpoll
)

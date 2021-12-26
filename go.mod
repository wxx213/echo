module github.wxx.example

go 1.15

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/kitex v0.0.8
	github.com/cloudwego/netpoll v0.1.2
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

// use customed packages
replace (
	github.com/cloudwego/kitex v0.0.8 => ../kitex
	github.com/cloudwego/netpoll v0.1.2 => ../netpoll
)

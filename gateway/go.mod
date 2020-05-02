module github.com/schmurfy/concourse-test/gateway

go 1.14

require (
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/schmurfy/concourse-test/proto v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/net v0.0.0-20190923162816-aa69164e4478 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.29.1
)

replace github.com/schmurfy/concourse-test/proto => ../proto

module github.com/schmurfy/concourse-test/service

go 1.14

require (
	github.com/blendle/zapdriver v1.3.1
	github.com/go-redis/redis/v7 v7.2.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/schmurfy/concourse-test/proto v0.0.0-00010101000000-000000000000
	github.com/schmurfy/test-lib2 v0.0.0-20200521105608-1e472016f9bf
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.29.1
)

replace github.com/schmurfy/concourse-test/proto => ../proto

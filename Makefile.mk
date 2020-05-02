# find out which version we are using and get their paths
MODELS_PATH=$(shell go list -m -json all | jq -r 'select(.Path == "github.com/schmurfy/concourse-test/proto") | .Dir' )
GOGOBUFPATH=$(shell go list -m -json all | jq -r 'select(.Path == "github.com/gogo/protobuf") | .Dir' )
GOLANGBUFPATH=$(shell go list -m -json all | jq -r 'select(.Path == "github.com/golang/protobuf") | .Dir' )

PROTODIRS 	:= $(shell ls $(MODELS_PATH))
PROTOS_IN   := $(shell find $(MODELS_PATH) -name *.proto)
PROTOS_OUT  := $(patsubst $(MODELS_PATH)/%.proto,generated_pb/%.pb.go,$(PROTOS_IN))

SERVICE 		:= $(shell basename $$PWD)

debug:
	@echo Paths:
	@echo MODELS_PATH: $(MODELS_PATH)
	@echo GOGOBUFPATH: $(GOGOBUFPATH)
	@echo GOLANGBUFPATH: $(GOLANGBUFPATH)
	@echo
	@echo Inputs:
	@echo $(PROTOS_IN)
	@echo
	@echo Outputs:
	@echo $(PROTOS_OUT)


generated_pb/%.pb.go: $(MODELS_PATH)/%.proto
	@echo Rebuilding $@...
	@mkdir -p $$(dirname $@)
	
	@protoc \
		--gogo_out=plugins=grpc:$$(dirname $@) \
		--proto_path=$$(dirname $<) \
		--proto_path=$(GOGOBUFPATH) \
		--proto_path=$(GOLANGBUFPATH) \
		$<


proto: $(PROTOS_OUT)

${SERVICE}: proto
	@echo Building ${SERVICE}...
	@go build -o ${SERVICE} main.go
	@echo Success

${SERVICE}-linux: proto
	mkdir -p  build
	GOARCH=amd64 GOOS=linux go build -o build/${SERVICE} main.go

image: ${SERVICE}-linux
	docker build -t schmurfy/concourse-test-${SERVICE}:latest -f Dockerfile ./build

push-image: image
	docker push schmurfy/concourse-test-${SERVICE}:latest

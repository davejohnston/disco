BIN=./bin/
GEN_DIR=./internal

PROTOBUF_INCLUDES += -I${GOPATH}/src
PROTOBUF_INCLUDES += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF_INCLUDES += -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/

all: dirs protoc gateway swagger controller

clean:
	rm -rf $(BIN) $(GEN_DIR)/api

dirs: $(BIN) $(GEN_DIR)

$(BIN):
	mkdir -p $(BIN)

$(GEN_DIR):
	mkdir -p $(GEN_DIR)


controller: protoc dirs
	go build -o bin/controller cmd/controller/controller.go

PROTO_FILES += api/secret.proto
PROTO_FILES += api/resource.proto
PROTO_FILES += api/provider.proto
PROTO_FILES += api/job.proto
PROTO_FILES += api/command.proto


protoc: dirs
	protoc -I/usr/local/include -I. \
		${PROTOBUF_INCLUDES} \
 	 	--go_out=plugins=grpc:$(GEN_DIR) \
		$(PROTO_FILES)

gateway: dirs
	protoc -I/usr/local/include -I. \
		${PROTOBUF_INCLUDES} \
     	--grpc-gateway_out=logtostderr=true:$(GEN_DIR) \
		$(PROTO_FILES)

swagger: dirs
	protoc -I/usr/local/include -I. \
		${PROTOBUF_INCLUDES} \
	--swagger_out=logtostderr=true:$(GEN_DIR) \
		$(PROTO_FILES)


check:
	gofmt -s -w .
	golint ./...
	govet ./...

doc:
	open http://localhost:6060/pkg/github.com/davejohnston/disco/?m=all
	godoc -http=:6060




//go:generate mkdir -p ../internal/api

//go:generate protoc --proto_path=$GOPATH/src/github.com/davejohnston/disco/api --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:$GOPATH/src/github.com/davejohnston/disco/internal/api secret.proto resource.proto provider.proto job.proto command.proto

//go:generate protoc --proto_path=$GOPATH/src/github.com/davejohnston/disco/api --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:$GOPATH/src/github.com/davejohnston/disco/internal/api command.proto

//go:generate protoc --proto_path=$GOPATH/src/github.com/davejohnston/disco/api --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/davejohnston/disco/vendor/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:$GOPATH/src/github.com/davejohnston/disco/internal/api command.proto

package api

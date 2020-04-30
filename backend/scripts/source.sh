# lines
# protoc user.proto --go_out=plugins=grpc:.
# protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 user.proto
export ALL_PROXY=
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
#export GOPATH=$(pwd)/..
#export GOBIN=$(GOPATH)/bin
#protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto

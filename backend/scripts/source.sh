#go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#protoc lorem.proto --go_out=plugins=grpc:.
export ALL_PROXY=
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
#export GOPATH=$(pwd)/..
#export GOBIN=$(GOPATH)/bin
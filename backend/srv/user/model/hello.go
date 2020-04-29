package user

import (
	"context"
	"google.golang.org/protobuf/internal/impl"
	. "nihplod/srv/user/endpoints"
	"nihplod/srv/user/pb"
)

func EncodeGRPCHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(HelloRequest)
	return &pb.HelloRequest{
		Name: req.Name,
	}, nil
}

func DecodeGRPCHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HelloRequest)
	return HelloRequest{
		Name: req.Name,
	}, nil
}

func EncodeGRPCHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(HelloResponse)
	return &pb.HelloResponse{
		Message: resp.Message,
	}, nil
}

func DecodeGRPCHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.HelloResponse)
	return HelloResponse{
		Message: resp.Message,
	}, nil
}
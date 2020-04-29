package model

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/pb"
)

func EncodeGRPCHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoints.HelloRequest)
	return &pb.HelloRequest{
		Name: req.Name,
	}, nil
}

func DecodeGRPCHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HelloRequest)
	return endpoints.HelloRequest{
		Name: req.Name,
	}, nil
}

func EncodeGRPCHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.HelloResponse)
	return &pb.HelloResponse{
		Message: resp.Message,
	}, nil
}

func DecodeGRPCHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.HelloResponse)
	return endpoints.HelloResponse{
		Message: resp.Message,
	}, nil
}

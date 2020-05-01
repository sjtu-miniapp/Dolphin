package model

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
)

func EncodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoints.GetUserRequest)
	return &pb.GetUserRequest{
		Id: req.Id,
	}, nil
}

func DecodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetUserRequest)
	return endpoints.GetUserRequest{
		Id: req.Id,
	}, nil
}

func EncodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.GetUserResponse)
	return &pb.GetUserResponse{
		Name: resp.Name,
		Gender: resp.Gender,
	}, nil
}

func DecodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.GetUserResponse)
	return endpoints.GetUserResponse{
		Name: resp.Name,
		Gender: resp.Gender,
	}, nil
}

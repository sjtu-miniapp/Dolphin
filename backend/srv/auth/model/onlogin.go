package model

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
)

func EncodeOnLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoints.OnLoginRequest)
	return &pb.OnLoginRequest{
		Code: req.Code,
	}, nil
}

func DecodeOnLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.OnLoginRequest)
	return endpoints.OnLoginRequest{
		Code: req.Code,
	}, nil
}

func EncodeOnLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.OnLoginResponse)
	return &pb.OnLoginResponse{
		Id: resp.Id,
		Sid: resp.Sid,
	}, nil
}

func DecodeOnLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.OnLoginResponse)
	return endpoints.OnLoginResponse{
		Id: resp.Sid,
	}, nil
}

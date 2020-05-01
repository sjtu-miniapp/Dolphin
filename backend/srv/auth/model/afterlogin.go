package model

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
)

func EncodeAfterLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoints.AfterLoginRequest)
	return &pb.AfterLoginRequest{
		Id: req.Id,
		Sid: req.Sid,
		Name: req.Name,
		Gender: req.Gender,
	}, nil
}

func DecodeAfterLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.AfterLoginRequest)
	return endpoints.AfterLoginRequest{
		Id: req.Id,
		Sid: req.Sid,
		Name: req.Name,
		Gender: req.Gender,
	}, nil
}

func EncodeAfterLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.AfterLoginResponse)
	return &pb.AfterLoginResponse{
		Err: int32(resp.Err),
	}, nil
}

func DecodeAfterLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.AfterLoginResponse)
	return endpoints.AfterLoginResponse{
		Err: int(resp.Err),
	}, nil
}

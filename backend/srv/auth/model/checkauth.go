package model

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
)

func EncodeCheckAuthRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoints.CheckAuthRequest)
	return &pb.CheckAuthRequest{
		Id: req.Id,
		Sid: req.Sid,
	}, nil
}

func DecodeCheckAuthnRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CheckAuthRequest)
	return endpoints.CheckAuthRequest{
		Id: req.Id,
		Sid: req.Sid,
	}, nil
}

func EncodeCheckAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.CheckAuthResponse)
	return &pb.CheckAuthResponse{
		Ok: resp.Ok,
	}, nil
}

func DecodeCheckAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.CheckAuthResponse)
	return endpoints.CheckAuthResponse{
		Ok: resp.Ok,
	}, nil
}

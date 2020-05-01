package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sjtu-miniapp/dolphin/auth/service"
)

type CheckAuthRequest struct {
	Id  string `json:"id"`
	Sid string `json:"sid"`
}

type CheckAuthResponse struct {
	Ok bool `json:"ok"`
}

func MakeCheckAuthEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckAuthRequest)
		ok, err := srv.CheckAuth(ctx, req.Id, req.Sid)
		if err != nil {
			return false, err
		}
		return CheckAuthResponse{Ok: ok}, nil
	}
}

func (e Endpoints) CheckAuth(ctx context.Context, id, sid string) (bool, error) {
	req := CheckAuthRequest{
		Id: id,
		Sid: sid,
	}
	resp, err := e.CheckAuthEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	checkAuthResp := resp.(CheckAuthResponse)
	return checkAuthResp.Ok, nil
}

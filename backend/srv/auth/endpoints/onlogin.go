package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sjtu-miniapp/dolphin/auth/service"
)

type OnLoginRequest struct {
	Code   string `json:"code"`
}

type OnLoginResponse struct {
	Id 		string `json:"id"`
	Sid		string `json:"sid"`
}

func MakeOnLoginEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OnLoginRequest)
		id, sid, err := srv.OnLogin(ctx, req.Code)
		if err != nil {
			return nil, err
		}
		return OnLoginResponse{Id: id, Sid: sid}, nil
	}
}

func (e Endpoints) OnLogin(ctx context.Context, code string) (string, string, error) {
	req := OnLoginRequest{
		Code: code,
	}
	resp, err := e.OnLoginEndpoint(ctx, req)
	if err != nil {
		return "", "", err
	}

	onLoginResp := resp.(OnLoginResponse)
	return onLoginResp.Id, onLoginResp.Sid, nil
}

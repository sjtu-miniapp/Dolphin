package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sjtu-miniapp/dolphin/auth/service"
)

type AfterLoginRequest struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender uint32    `json:"gender"`
	Sid    string `json:"sid"`
}

type AfterLoginResponse struct {
	Err int `json:"err"`
}

func MakeAfterLoginEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AfterLoginRequest)
		_err, err := srv.AfterLogin(ctx, req.Id, req.Name, req.Gender, req.Sid)
		if err != nil {
			return 1, err
		}
		return AfterLoginResponse{Err: _err}, nil
	}
}

func (e Endpoints) AfterLogin(ctx context.Context, id, name string, gender uint32, sid string) (int, error) {
	req := AfterLoginRequest{
		Id: id,
		Name: name,
		Gender: gender,
		Sid: sid,
	}
	resp, err := e.AfterLoginEndpoint(ctx, req)
	if err != nil {
		return 1, err
	}

	afterLoginResp := resp.(AfterLoginResponse)
	return afterLoginResp.Err, nil
}

package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sjtu-miniapp/dolphin/auth/service"
)

type GetUserRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct {
	Name  string `json:"name"`
	Gender uint32 `json:"gender"`
}

func MakeGetUserEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		name, gender, err := srv.GetUser(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return GetUserResponse{Name: name, Gender: gender}, nil
	}
}

func (e Endpoints) GetUser(ctx context.Context, id string) (string, uint32, error) {
	req := GetUserRequest{
		Id: id,
	}
	resp, err := e.GetUserEndpoint(ctx, req)
	if err != nil {
		return "", 0, err
	}
	getUserResp := resp.(GetUserResponse)
	return getUserResp.Name, getUserResp.Gender, nil
}

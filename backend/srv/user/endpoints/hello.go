package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sjtu-miniapp/dolphin/user/service"
)

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Message string `json:"message"`
}

func MakeHelloEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		message, err := srv.Hello(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return HelloResponse{Message: message}, nil
	}
}

func (e Endpoints) Hello(ctx context.Context, name string) (string, error) {
	req := HelloRequest{
		Name: name,
	}
	resp, err := e.HelloEndpoint(ctx, req)
	if err != nil {
		return "", err
	}

	helloResp := resp.(HelloResponse)
	return helloResp.Message, nil
}
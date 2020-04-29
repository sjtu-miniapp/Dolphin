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
	Message string
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
package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	service "nihplod/srv/user/service"
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
		message, err := srv.Hello(ctx, req.name)

		if err != nil {
			return nil, err
		}
		return HelloResponse{message: message}, nil
	}
}
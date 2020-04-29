package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	. "nihplod/srv/user/model"
	"nihplod/srv/user/pb"
	endpoints "nihplod/srv/user/endpoints"
)

type grpcServer struct {
	hello grpctransport.Handler
}

// service
func NewGRPCServer(_ context.Context, endpoint endpoints.Endpoints) pb.UserServer {
	return &grpcServer{
		hello: grpctransport.NewServer(
			endpoint.HelloEndpoint,
			DecodeGRPCHelloRequest,
			EncodeGRPCHelloResponse,
			),
	}
}
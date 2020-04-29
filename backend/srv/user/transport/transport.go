package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sjtu-miniapp/dolphin/user/model"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/pb"
)

type grpcServer struct {
	hello grpctransport.Handler
}

// service
func NewGRPCServer(_ context.Context, endpoint endpoints.Endpoints) pb.UserServer {
	return &grpcServer{
		hello: grpctransport.NewServer(
			endpoint.HelloEndpoint,
			model.DecodeGRPCHelloRequest,
			model.EncodeGRPCHelloResponse,
			),
	}
}
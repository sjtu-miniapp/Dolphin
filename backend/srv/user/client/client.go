package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/model"
	"github.com/sjtu-miniapp/dolphin/user/pb"
	"github.com/sjtu-miniapp/dolphin/user/service"
	"google.golang.org/grpc"
)

// Return new user service
func New(conn *grpc.ClientConn) service.Service {
	var helloEndpoint = grpctransport.NewClient(
		conn,
		"pb.UserService",
		"Hello",
		model.EncodeGRPCHelloRequest,
		model.DecodeGRPCHelloResponse,
		pb.HelloResponse{},
	).Endpoint()
	return endpoints.Endpoints{
		HelloEndpoint: helloEndpoint,
	}
}

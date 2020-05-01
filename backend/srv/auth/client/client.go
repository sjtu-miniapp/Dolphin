package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/model"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
	"github.com/sjtu-miniapp/dolphin/auth/service"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) service.Service {
	var onLoginEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"OnLogin",
		model.EncodeOnLoginRequest,
		model.DecodeOnLoginResponse,
		pb.OnLoginResponse{},
	).Endpoint()
	var afterLoginEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"AfterLogin",
		model.EncodeAfterLoginRequest,
		model.DecodeAfterLoginResponse,
		pb.AfterLoginResponse{},
	).Endpoint()
	var getUserEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"GetUser",
		model.EncodeGetUserRequest,
		model.DecodeGetUserResponse,
		pb.GetUserResponse{},
	).Endpoint()

	var checkAuthEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"CheckAuth",
		model.EncodeCheckAuthRequest,
		model.DecodeCheckAuthnRequest,
		pb.CheckAuthResponse{},
	).Endpoint()

	return endpoints.Endpoints{
		OnLoginEndpoint:    onLoginEndpoint,
		AfterLoginEndpoint: afterLoginEndpoint,
		GetUserEndpoint:    getUserEndpoint,
		CheckAuthEndpoint:  checkAuthEndpoint,
	}
}

package transport

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/model"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
	"golang.org/x/net/context"
)

type grpcServer struct {
	onLogin    grpctransport.Handler
	getUser    grpctransport.Handler
	afterLogin grpctransport.Handler
	checkAuth  grpctransport.Handler
}

func (s *grpcServer) OnLogin(ctx context.Context, request *pb.OnLoginRequest) (*pb.OnLoginResponse, error) {
	_, resp, err := s.onLogin.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.OnLoginResponse), nil
}

func (s *grpcServer) AfterLogin(ctx context.Context, request *pb.AfterLoginRequest) (*pb.AfterLoginResponse, error) {
	_, resp, err := s.onLogin.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AfterLoginResponse), nil
}

func (s *grpcServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUserResponse), nil
}

func (s *grpcServer) CheckAuth(ctx context.Context, request *pb.CheckAuthRequest) (*pb.CheckAuthResponse, error) {
	_, resp, err := s.checkAuth.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CheckAuthResponse), nil
}

// service
func NewGRPCServer(_ context.Context, endpoint endpoints.Endpoints) pb.AuthServiceServer {
	return &grpcServer{
		// no context
		onLogin: grpctransport.NewServer(
			endpoint.OnLoginEndpoint,
			model.DecodeOnLoginRequest,
			model.EncodeOnLoginResponse,
		),
		afterLogin: grpctransport.NewServer(
			endpoint.AfterLoginEndpoint,
			model.DecodeAfterLoginRequest,
			model.DecodeAfterLoginResponse,
		),
		getUser: grpctransport.NewServer(
			endpoint.GetUserEndpoint,
			model.DecodeGetUserRequest,
			model.EncodeGetUserResponse,
		),
		checkAuth: grpctransport.NewServer(
			endpoint.CheckAuthEndpoint,
			model.DecodeCheckAuthnRequest,
			model.EncodeCheckAuthResponse,
		),
	}
}

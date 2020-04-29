package user

import (
	"context"
	"nihplod/srv/user/pb"
)

func (s *grpcServer) Hello(ctx context.Context, r*pb.HelloRequest) (*pb.HelloResponse, error) {
	_, resp, err := s.hello.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HelloResponse), nil
}

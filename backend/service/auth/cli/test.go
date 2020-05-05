package main

import (
	"context"
	"fmt"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/sjtu-miniapp/dolphin/utils/parse"
)

type Config struct {
	Registry []string `yaml:"registry"`
}

func main() {
	var cfg Config
	err := parse.LoadConfig(&cfg)
	if err != nil {
		return
	}
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = cfg.Registry
	})
	service := micro.NewService(
		micro.Name("go.micro.cli.auth"),
		micro.Registry(reg),
		micro.Flags(
			&cli.StringFlag{
				Name:  "cfg",
				Usage: "location of config file",
			},
		),
	)
	service.Init()
	auth := pb.NewAuthService("go.micro.srv.auth", service.Client())

	resp, err := auth.OnLogin(context.TODO(), &pb.OnLoginRequest{
		Code: "abcdef",
	})
	fmt.Println(err)
	fmt.Println(resp)

	r, err := auth.CheckAuth(context.TODO(), &pb.CheckAuthRequest{
		Openid: resp.Openid,
		Sid:    resp.Sid,
	})
	fmt.Println(r)
	_, err = auth.PutUser(context.Background(), &pb.PutUserRequest{
		Openid: resp.Openid,
		Name:   "hello",
		Gender: "F",
		Avatar: "world",
	})
	fmt.Println(err)
	rs, _ := auth.GetUser(context.Background(), &pb.GetUserRequest{
		Id: resp.Openid,
	})
	fmt.Println(rs)
	if err != nil {
		fmt.Println(err)
	}
}

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
	auth := pb.NewAuthService("go.micro.srv.group", service.Client())

	//rsp, err := auth.GetAuth(context.TODO(), &pb.GetAuthRequest{
	//	Id: 1,
	//})
	//fmt.Println(rsp)
	if err != nil {
		fmt.Println(err)
	}
}

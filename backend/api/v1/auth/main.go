package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/debug/log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/sjtu-miniapp/dolphin/utils/parse"
	"os"
	"os/signal"
	"syscall"
)

var (
	srv pb.AuthService
)

type Config struct {
	Version  string   `yaml:"version"`
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
	var service web.Service
	{
		service = web.NewService(
			web.Registry(reg),
			web.Name("go.micro.api.auth"),
			web.Flags(
				&cli.StringFlag{
					Name:  "cfg",
					Usage: "location of config file",
				},
			),
		)
		_ = service.Init()
		srv = pb.NewAuthService("go.micro.srv.auth", client.DefaultClient)
		base := "/api/" + cfg.Version
		router := Router(base)
		service.Handle("/", router)
	}

	errChan := make(chan error)
	go func() {
		if err := service.Run(); err != nil {
			log.Fatal("fail to run the service", err)
			errChan <- err
			return
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	log.Error("shutting down the service...", <-errChan)
}

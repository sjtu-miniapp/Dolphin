package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/sjtu-miniapp/dolphin/service/auth/srv/impl"
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/utils/parse"
	"time"
)

func createService(c Config) micro.Service {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = c.Registry
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Flags(
			&cli.StringFlag{
				Name:  "cfg",
				Usage: "location of config file",
			},
		),
	)
	service.Init()
	return service
}

type Config struct {
	Mysql struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Db   string `yaml:"db"`
	} `yaml:"mysql"`
	Registry []string `yaml:"registry"`
	App struct {
		AppId string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	} `yaml:"app"`
	Redis struct{
		Host string `yaml:"host"`
		Pass string `yaml:"pass"`
	} `yaml:"redis"`
	Debug int `yaml:"debug"`
}

func main() {
	var cfg Config
	err := parse.LoadConfig(&cfg)
	if err != nil {
		return
	}

	srv := createService(cfg)
	sqldb, err := database.DbConn(cfg.Mysql.User, cfg.Mysql.Pass,
		cfg.Mysql.Host, cfg.Mysql.Db, 3306, cfg.Debug)
	if err != nil {
		log.Fatal(err)
		return
	}
	redisdb, err := database.InitRedis(cfg.Redis.Host, cfg.Redis.Pass)
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = pb.RegisterAuthHandler(srv.Server(), &impl.Auth{
		SqlDb:     sqldb,
		RedisDb:   redisdb,
		AppId:     cfg.App.AppId,
		AppSecret: cfg.App.AppSecret,
	})
	if err := srv.Run(); err != nil {
		log.Fatal("fail to run the service", err)
	}
}

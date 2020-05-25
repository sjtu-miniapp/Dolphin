package main

import (
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"github.com/sjtu-miniapp/dolphin/service/task/srv/impl"
	"github.com/sjtu-miniapp/dolphin/utils/parse"
	"time"
)

func createService(c Config) micro.Service {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = c.Registry
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.task"),
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
	Mongo struct{
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Port int `yaml:"port"`
		Db   string `yaml:"db"`
	} `yaml:"mongo"`
	Registry []string `yaml:"registry"`
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
	mongo, err := database.MongoConn(cfg.Mongo.Host, cfg.Mongo.User, cfg.Mongo.Pass, cfg.Mongo.Port)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer mongo.Disconnect(context.TODO())
	mongoDb := mongo.Database(cfg.Mongo.Db)
	_ = pb.RegisterTaskHandler(srv.Server(), &impl.Task{
		SqlDb:   sqldb,
		MongoDb: mongoDb,
	})
	if err := srv.Run(); err != nil {
		log.Fatal("fail to run the service", err)
	}
}
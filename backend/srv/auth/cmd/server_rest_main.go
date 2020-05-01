package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sjtu-miniapp/dolphin/auth/endpoints"
	"github.com/sjtu-miniapp/dolphin/auth/logger"
	"github.com/sjtu-miniapp/dolphin/auth/pb"
	"github.com/sjtu-miniapp/dolphin/auth/rest"
	"github.com/sjtu-miniapp/dolphin/auth/service"
	"github.com/sjtu-miniapp/dolphin/auth/transport"
	"go.uber.org/zap"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	AuthServer struct {
		// for other services to use
		Host     string `yaml:"host"`
		HttpPort string `yaml:"httpPort"`
		GrpcPort string `yaml:"grpcPort"`
	} `yaml:"authServer"`
	Sqldb struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string `yaml:"host"`
		Db       string `yamle:"db"`
	} `yaml:"mysql"`
	Log struct {
		LogLevel      int    `yaml:"logLevel"`
		LogTimeFormat string `yaml:"logTimeFormat"`
	} `yaml:"log"`
	App struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	} `yaml:"app"`
}

func main() {
	cfgfile := flag.String("cfg", "", "config file")
	flag.Parse()
	var cfg Config
	err := readFile(&cfg, *cfgfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// initialize logger
	if err := logger.Init(cfg.Log.LogLevel, cfg.Log.LogTimeFormat); err != nil {
		return
	}

	ctx := context.Background()
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.Sqldb.Username,
		cfg.Sqldb.Password,
		cfg.Sqldb.Host,
		cfg.Sqldb.Db,
		param,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Log.Fatal("failed to open database:", zap.String("reason", err.Error()))
	}
	err = db.Ping()
	//err = db.Ping()
	errChan := make(chan error)
	if err != nil {
		logger.Log.Fatal("failed to connect to database:", zap.String("reason", err.Error()))
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Log.Fatal("failed to close database:", zap.String("reason", err.Error()))
		}
	}()

	// init service
	var svc service.Service
	svc = service.AuthService{
		Db: db,
		AppId: cfg.App.AppId,
		AppSecret: cfg.App.AppSecret,
	}

	// creating Endpoints struct
	epts := endpoints.Endpoints{
		OnLoginEndpoint:    endpoints.MakeOnLoginEndpoint(svc),
		AfterLoginEndpoint: endpoints.MakeAfterLoginEndpoint(svc),
		GetUserEndpoint:    endpoints.MakeGetUserEndpoint(svc),
		CheckAuthEndpoint:  endpoints.MakeCheckAuthEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", ":"+cfg.AuthServer.GrpcPort)
		if err != nil {
			errChan <- err
			return
		}
		handler := transport.NewGRPCServer(ctx, epts)
		gRPCServer := grpc.NewServer()
		pb.RegisterAuthServiceServer(gRPCServer, handler)
		logger.Log.Info("starting gRPC server...")
		errChan <- gRPCServer.Serve(listener)
	}()

	// run HTTP gateway
	go func() {
		err := rest.RunServer(ctx, cfg.AuthServer.GrpcPort, cfg.AuthServer.HttpPort)
		if err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
		logger.Log.Warn("shutting down gRPC server...")
	}()

	fmt.Println(<-errChan)
}

func readFile(cfg *Config, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	return err
}

package main

import (
	"flag"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/pb"
	"github.com/sjtu-miniapp/dolphin/user/rest"
	"github.com/sjtu-miniapp/dolphin/user/service"
	"github.com/sjtu-miniapp/dolphin/user/transport"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	//DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	//DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	//DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	//DatastoreDBSchema string
}

func main() {
	//var gRPCAddr = flag.String("grpc", ":8081",
	//		"gRPC listen address")
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "8081", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "9081", "HTTP port to bind")
	//flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	//flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	//flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	//flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()
	ctx := context.Background()

	// init service
	var svc service.Service
	svc = service.UserService{}
	errChan := make(chan error)

	// creating Endpoints struct
	epts := endpoints.Endpoints{
		HelloEndpoint: endpoints.MakeHelloEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", ":8081")
		if err != nil {
			errChan <- err
			return
		}
		handler := transport.NewGRPCServer(ctx, epts)
		gRPCServer := grpc.NewServer()
		pb.RegisterUserServiceServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	// run HTTP gateway
	go func() {
		err := rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
		if err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}

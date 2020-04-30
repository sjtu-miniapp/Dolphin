package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/pb"
	"github.com/sjtu-miniapp/dolphin/user/rest"
	"github.com/sjtu-miniapp/dolphin/user/service"
	"github.com/sjtu-miniapp/dolphin/user/transport"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort    string
	SQLHost     string
	SQLUser     string
	SQLPassword string
	SQLDatabase string
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "8081", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "9081", "HTTP port to bind")
	flag.StringVar(&cfg.SQLHost, "sql-host", "", "sql host")
	flag.StringVar(&cfg.SQLUser, "sql-user", "", "sql user")
	flag.StringVar(&cfg.SQLPassword, "sql-passwd", "", "sql password")
	flag.StringVar(&cfg.SQLDatabase, "sql-db", "", "sql database")
	flag.Parse()
	ctx := context.Background()

	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.SQLUser,
		cfg.SQLPassword,
		cfg.SQLHost,
		cfg.SQLDatabase,
		param)
	db, _ := sql.Open("mysql", dsn)
	db, _ = sql.Open("mysql",
		"root:610878@tcp(127.0.0.1:3306)/test")
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// init service
	var svc service.Service
	svc = service.UserService{
		Db: db,
	}
	errChan := make(chan error)

	// creating Endpoints struct
	epts := endpoints.Endpoints{
		HelloEndpoint: endpoints.MakeHelloEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
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

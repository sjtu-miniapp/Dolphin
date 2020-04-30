package main

import (
	"context"
	"flag"
	grpcClient "github.com/sjtu-miniapp/dolphin/user/client"
	"github.com/sjtu-miniapp/dolphin/user/service"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "gRPC address")
	)

	flag.Parse()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx,
		*grpcAddr,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	log.Print("Connection is built.")
	defer conn.Close()

	userService := grpcClient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "hello":
		name, _ := pop(args)
		msg, err := hello(ctx, userService, name)
		if err != nil {
			log.Print(err)
		}
		log.Print(msg)
	default:
		log.Fatalln("Bad method", cmd)
	}

}

// parse command line argument one by one
func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// call hello service
func hello(ctx context.Context, srv service.Service, name string) (string, error) {
	msg, err := srv.Hello(ctx, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return msg, err
}

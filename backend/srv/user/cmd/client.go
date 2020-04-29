package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
	grpcClient "github.com/ru-rocker/gokit-playground/lorem-grpc/client"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "gRPC address")
	)

	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx,
		*grpcAddr,
		grpc.WithInsecure(),
		context.WithTimeout()
	)

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	userService := grpcClient.New(conn)

}
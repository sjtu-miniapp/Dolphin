package main
import (
	"flag"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/user/endpoints"
	"github.com/sjtu-miniapp/dolphin/user/pb"
	"github.com/sjtu-miniapp/dolphin/user/service"
	"github.com/sjtu-miniapp/dolphin/user/transport"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		gRPCAddr = flag.String("grpc", ":8081",
			"gRPC listen address")
	)

	flag.Parse()
	ctx := context.Background()

	// init  service
	var svc service.Service
	svc = service.UserService{}
	errChan := make(chan error)

	// creating Endpoints struct
	epts := endpoints.Endpoints{
		HelloEndpoint: endpoints.MakeHelloEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := transport.NewGRPCServer(ctx, epts)
		gRPCServer := grpc.NewServer()
		pb.RegisterUserServiceServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()


	fmt.Println(<- errChan)
}
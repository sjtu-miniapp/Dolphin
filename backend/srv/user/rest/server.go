package rest


import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/sjtu-miniapp/dolphin/user/pb"
)

// RunServer runs HTTP/REST gateway
// TODO: support https
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:" + grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
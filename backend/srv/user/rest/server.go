package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sjtu-miniapp/dolphin/user/logger"
	"github.com/sjtu-miniapp/dolphin/user/middleware"
	"github.com/sjtu-miniapp/dolphin/user/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
// TODO: support https
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}
	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: middleware.AddRequestID(
			middleware.AddLogger(logger.Log,
				middleware.CORS(mux))),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

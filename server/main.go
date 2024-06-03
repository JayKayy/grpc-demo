package main

import (
	"log/slog"
	"net"

	"github.com/grpc-demo/app"
	"github.com/grpc-demo/internal/config"
	"github.com/grpc-demo/internal/server"
	"google.golang.org/grpc"
)

func main() {
	srv := server.New("one")
	lis, err := net.Listen("tcp", config.SrvAddr)
	if err != nil {
		slog.Error("listening: %w", err)
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	app.RegisterMessengerServer(grpcServer, srv)
	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("grpc server listening", "err", err)
	}
}

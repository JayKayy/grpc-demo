package main

import (
	"log/slog"
	"net"
	"time"

	"github.com/grpc-demo/app"
	"github.com/grpc-demo/client"
	"github.com/grpc-demo/server"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "8080"
)

var srvAddr = net.JoinHostPort(host, port)

func main() {
	srv := server.New("one")
	c, err := client.New(srvAddr)
	if err != nil {
		slog.Error("creating client: %w", err)
		return
	}
	lis, err := net.Listen("tcp", srvAddr)
	if err != nil {
		slog.Error("listening: %w", err)
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	app.RegisterMessengerServer(grpcServer, srv)
	go grpcServer.Serve(lis)
	err = c.SendMessage("hello")
	if err != nil {
		slog.Error("sending message", "err", err)
	}
	time.Sleep(2 * time.Second)
}

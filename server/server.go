package server

import (
	"context"
	"log/slog"

	"github.com/grpc-demo/app"
)

func New(name string) *Server {
	return &Server{
		Name: name,
	}
}

type Server struct {
	Name string
	app.UnimplementedMessengerServer
}

func (s Server) SendMessage(ctx context.Context, msg *app.Msg) (*app.Msg, error) {
	slog.Info("got message", "msg", msg.Content)

	return nil, nil
}

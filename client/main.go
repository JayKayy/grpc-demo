package main

import (
	"log/slog"

	"github.com/grpc-demo/internal/client"
	"github.com/grpc-demo/internal/config"
)

func main() {
	c, err := client.New(config.SrvAddr)
	if err != nil {
		slog.Error("creating client: %w", err)
		return
	}
	err = c.SendMessage("hello")
	if err != nil {
		slog.Error("sending message", "err", err)
	}
}

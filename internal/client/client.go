package client

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-demo/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Messenger app.MessengerClient
}

func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connecting to server: %w", err)
	}
	client := app.NewMessengerClient(conn)
	return &Client{Messenger: client}, nil
}

func (c Client) SendMessage(msg string) error {
	ctxBg := context.Background()
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	m := &app.Msg{
		Content: msg,
	}
	_, err := c.Messenger.SendMessage(ctx, m)
	if err != nil {
		return fmt.Errorf("send failed: %w", err)
	}
	return nil
}

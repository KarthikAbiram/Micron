package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// Client wraps the gRPC connection and generated client
type Client struct {
	conn   *grpc.ClientConn
	client MicronGRPCClient
}

func New(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := NewMicronGRPCClient(conn) // generated client
	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) SendMessage(command, payload string) (*MessageReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &MessageRequest{
		Command: command,
		Payload: payload,
	}

	resp, err := c.client.Message(ctx, req) // "Message" RPC from proto
	if err != nil {
		return nil, err
	}

	return resp, nil
}

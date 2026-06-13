package main

import (
	"context"
	"log"
	"time"

	pb "micron-template/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMicronGRPCClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	resp, err := client.Message(
		ctx,
		&pb.MessageRequest{
			Command: "PING",
			Payload: "Hello Server",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Payload: %s", resp.Payload)
	log.Printf(
		"Status: code=%d error=%v desc=%s",
		resp.Status.Code,
		resp.Status.IsError,
		resp.Status.Desc,
	)
}

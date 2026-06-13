package main

import (
	"context"
	"log"
	"net"

	pb "micron-template/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMicronGRPCServer
}

func (s *server) Message(
	ctx context.Context,
	req *pb.MessageRequest,
) (*pb.MessageReply, error) {

	log.Printf("Received command=%s payload=%s",
		req.Command,
		req.Payload,
	)

	return &pb.MessageReply{
		Payload: "Processed: " + req.Payload,
		Status: &pb.MessageReply_Status{
			Code:    0,
			IsError: false,
			Desc:    "Success",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMicronGRPCServer(grpcServer, &server{})

	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

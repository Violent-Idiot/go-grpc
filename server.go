package main

import (
	"grpc/chat"
	// "grpc/proto"
	"log"
	"net"

	// "github.com/golang/protobuf/protoc-gen-go/grpc"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/encoding/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	// proto.RegisterChatServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server %v", err)
	}

}

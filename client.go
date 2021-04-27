package main

import (
	"context"
	"grpc/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server Failed: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	msg := chat.Message{
		Body: "Yo!",
	}

	res, err := c.SayHello(context.Background(), &msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Res from Server: %s", res.Body)
}

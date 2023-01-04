package main

import (
	"context"
	message "github.com/arana-db/arana/grpc/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := message.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &message.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Fatalf("greet response: %v", resp)
}

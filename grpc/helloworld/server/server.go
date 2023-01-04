package main

import (
	message "github.com/arana-db/arana/grpc/helloworld/proto"
	"github.com/arana-db/arana/grpc/helloworld/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	srv := grpc.NewServer()
	message.RegisterMessageSenderServer(srv, &service.MessageSenderService{message.UnimplementedMessageSenderServer{}, &service.SenderService{}})

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

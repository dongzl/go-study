package service

import (
	"context"
	message "github.com/arana-db/arana/grpc/helloworld/proto"
	"log"
)

type MessageSenderService struct {
	message.UnimplementedMessageSenderServer
	Service *SenderService
}

func (m MessageSenderService) Send(ctx context.Context, request *message.MessageRequest) (*message.MessageResponse, error) {
	return m.Service.handleSendMessage(ctx, request)
}

type SenderService struct {
}

func (s SenderService) handleSendMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	log.Println("receive message:", req.GetSaySomething())
	resp := &message.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}

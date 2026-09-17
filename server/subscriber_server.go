package server

import (
	"context"

	subscriber "github.com/yoshihio/go-memps/pkg/subscriber/proto"
	"github.com/yoshihio/go-memps/service"
)

type SubscriberServer struct {
	subscriber.UnimplementedSubscriberServiceServer
	QueueService service.QueueService
}

func NewSubscriberServer(q service.QueueService) SubscriberServer {
	return SubscriberServer{QueueService: q}
}

func (s SubscriberServer) Subscribe(ctx context.Context, req *subscriber.SubscribeRequest) (*subscriber.SubscribeResponse, error) {
	return &subscriber.SubscribeResponse{
		Message: s.QueueService.Dequeue().Topic,
	}, nil
}

func (s SubscriberServer) SubscribeBytes(ctx context.Context, req *subscriber.SubscribeBytesRequest) (*subscriber.SubscribeBytesResponse, error) {
	return &subscriber.SubscribeBytesResponse{}, nil
}

package server

import (
	"context"

	"github.com/yoshihio/go-memps/entity"
	publisher "github.com/yoshihio/go-memps/pkg/publisher/proto"
	"github.com/yoshihio/go-memps/service"
)

type PublisherServer struct {
	publisher.UnimplementedPublisherServiceServer
	queue service.QueueService
}

func NewPublisherServer(queue service.QueueService) PublisherServer {
	return PublisherServer{queue: queue}
}

func (s PublisherServer) Publish(ctx context.Context, req *publisher.PublishRequest) (*publisher.PublishResponse, error) {
	s.queue.Enqueue(entity.NewData(req.Topic, req.Message))
	return &publisher.PublishResponse{}, nil
}

func (s PublisherServer) PublishBytes(ctx context.Context, req *publisher.PublishBytesRequest) (*publisher.PublishBytesResponse, error) {
	return &publisher.PublishBytesResponse{}, nil
}

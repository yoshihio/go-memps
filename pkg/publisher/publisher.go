package grpc_client

import (
	"context"

	publisher "github.com/yoshihio/go-memps/pkg/publisher/proto"
	"google.golang.org/grpc"
)

type PublisherService interface {
	Publish(ctx context.Context, in *publisher.PublishRequest, opts ...grpc.CallOption) (*publisher.PublishResponse, error)
	PublishBytes(ctx context.Context, in *publisher.PublishBytesRequest, opts ...grpc.CallOption) (*publisher.PublishBytesResponse, error)
}

type PublisherServiceImpl struct {
	publisher.PublisherServiceClient
}

func NewPublisherService(client publisher.PublisherServiceClient) PublisherService {
	return &PublisherServiceImpl{client}
}

func (s *PublisherServiceImpl) Publish(ctx context.Context, in *publisher.PublishRequest, opts ...grpc.CallOption) (*publisher.PublishResponse, error) {
	return s.PublisherServiceClient.Publish(ctx, in, opts...)
}

func (s *PublisherServiceImpl) PublishBytes(ctx context.Context, in *publisher.PublishBytesRequest, opts ...grpc.CallOption) (*publisher.PublishBytesResponse, error) {
	return s.PublisherServiceClient.PublishBytes(ctx, in, opts...)
}

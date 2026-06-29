package grpc_client

import (
	"context"

	publisher "github.com/yoshihio/go-memps/proto/go-memps-proto/publisher/v1"
	"google.golang.org/grpc"
)

type PublisherService struct {
	publisher.PublisherServiceClient
}

func NewPublisherService(client publisher.PublisherServiceClient) *PublisherService {
	return &PublisherService{client}
}

func (s *PublisherService) Publish(ctx context.Context, in *publisher.PublishRequest, opts ...grpc.CallOption) (*publisher.PublishResponse, error) {
	return s.PublisherServiceClient.Publish(ctx, in, opts...)
}

func (s *PublisherService) PublishBytes(ctx context.Context, in *publisher.PublishBytesRequest, opts ...grpc.CallOption) (*publisher.PublishBytesResponse, error) {
	return s.PublisherServiceClient.PublishBytes(ctx, in, opts...)
}

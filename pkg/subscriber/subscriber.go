package grpc_client

import (
	subscriber "github.com/yoshihio/go-memps/pkg/subscriber/proto"
)

type SubscriberService interface {
	Subscribe() error
}

func NewSubscriber(subscriber.SubscriberServiceClient) SubscriberService {
	return &SubscriberImpl{}
}

type SubscriberImpl struct {
}

func (s *SubscriberImpl) Subscribe() error {
	return nil
}

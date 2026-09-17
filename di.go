package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/yoshihio/go-memps/config"
	"github.com/yoshihio/go-memps/server"
	"github.com/yoshihio/go-memps/service"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	c := dig.New()

	c.Provide(initConfig)
	c.Provide(service.NewQueue)
	c.Provide(server.NewPublisherServer)
	c.Provide(server.NewSubscriberServer)

	// c.Provide(grpc_client.NewPublisherService)
	// c.Provide(publisher.NewPublisherServiceClient)
	// c.Provide(subscriber.NewSubscriberServiceClient)

	return c
}

func initConfig() config.Config {
	cfg := config.Config{}
	godotenv.Load()
	envconfig.MustProcess("", &cfg)
	return cfg
}

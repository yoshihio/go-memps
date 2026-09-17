package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	publisher "github.com/yoshihio/go-memps/pkg/publisher/proto"
	subscriber "github.com/yoshihio/go-memps/pkg/subscriber/proto"
	"github.com/yoshihio/go-memps/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting")
	defer log.Println("Exiting")
	c := NewContainer()

	closeChan := make(chan struct{})
	canShutDown := make(chan struct{})

	go func() {
		defer close(closeChan)
		err := c.Invoke(func(publisherServer server.PublisherServer, subscriberServer server.SubscriberServer) {
			opts := []grpc.ServerOption{
				grpc.ConnectionTimeout(100 * time.Millisecond),
				grpc.MaxRecvMsgSize(1 * 1024 * 1024),
			}
			gs := grpc.NewServer(opts...)
			publisher.RegisterPublisherServiceServer(gs, &publisherServer)
			subscriber.RegisterSubscriberServiceServer(gs, &subscriberServer)

			lis, err := net.Listen("tcp", ":8181")
			if err != nil {
				panic(err)
			}

			reflection.Register(gs)

			go func() {
				<-closeChan
				gs.GracefulStop()
				canShutDown <- struct{}{}
			}()

			log.Println("Server started on", lis.Addr())
			if err = gs.Serve(lis); err != nil {
				panic(err)
			}
		})
		if err != nil {
			panic(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sig:
			log.Println("Signal received, shutting down")
			closeChan <- struct{}{}
		case <-canShutDown:
			log.Println("Server stopped")
			return
		}
	}
}

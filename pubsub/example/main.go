package main

import (
	"context"
	"go-food-delivery/pubsub"
	"go-food-delivery/pubsub/pblocal"
	"log"
	"time"
)

func main() {
	var localPb = pblocal.NewLocalPubSub()

	var orderTopic pubsub.Topic = "OrderCreated"

	localPb.Publish(context.Background(), orderTopic, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), orderTopic, pubsub.NewMessage(2))
	localPb.Publish(context.Background(), orderTopic, pubsub.NewMessage(3))
	localPb.Publish(context.Background(), orderTopic, pubsub.NewMessage("abc"))

	sub1, close1 := localPb.Subscribe(context.Background(), orderTopic)
	sub2, _ := localPb.Subscribe(context.Background(), orderTopic)
	sub3, _ := localPb.Subscribe(context.Background(), orderTopic)

	go func() {
		for {
			log.Println("con 1 :", (<-sub1).Data())
			//time.Sleep(time.Millisecond * 800)
		}
	}()

	go func() {
		for {
			log.Println("con 2 :", (<-sub2).Data())
			//time.Sleep(time.Millisecond * 800)
		}
	}()

	go func() {
		for {
			log.Println("con 3 :", (<-sub3).Data())
			time.Sleep(time.Millisecond * 800)
		}
	}()

	time.Sleep(time.Second * 5)
	close1()

	localPb.Publish(context.Background(), orderTopic, pubsub.NewMessage("another topic"))

	time.Sleep(time.Second * 10)
}

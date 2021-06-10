package pblocal

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/pubsub"
	"log"
	"sync"
)

// A pub-sub run locally (in memory)
// It has a queue (buffer channel) and it's core and many group of subscribers
// Because of we want to send a message with specific topic for many subscribers in group can handle

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	p := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}
	p.run()

	return p
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetTopic(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published: ", data.String(), " with data ", data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("UnSubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}

}

// TODO: move it to engine
func (ps *localPubSub) run() error {
	log.Println("PubSub started")

	go func() {
		for {
			mess := <-ps.messageQueue

			if subs, ok := ps.mapChannel[mess.Topic()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
			//else {
			//	ps.messageQueue <- mess
			//}
		}
	}()

	return nil
}

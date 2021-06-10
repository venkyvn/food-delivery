package pubsub

import "context"

type Topic string

type PubSub interface {
	Publish(ctx context.Context, topic Topic, data *Message) error
	Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, close func())
	// UnSubscribe(ctx context.Context, topic Channel) error
}

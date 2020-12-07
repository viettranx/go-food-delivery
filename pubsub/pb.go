package pubsub

import (
	"context"
)

type Channel string

type Pubsub interface {
	Publish(ctx context.Context, channel Channel, data *Message) error
	Subscribe(ctx context.Context, channel Channel) (ch <-chan *Message, close func())
}

package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channel   Channel // can be ignore
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.Nanosecond()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channel)
}

func (evt *Message) Channel() Channel {
	return evt.channel
}

func (evt *Message) SetChannel(channel Channel) {
	evt.channel = channel
}

func (evt *Message) Data() interface{} {
	return evt.data
}

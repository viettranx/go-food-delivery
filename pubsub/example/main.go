package main

import (
	"context"
	"fooddlv/pubsub"
	"fooddlv/pubsub/pblocal"
	"log"
	"time"
)

func main() {
	var localPb pubsub.Pubsub = pblocal.NewPubSub()
	chn := pubsub.Channel("OrderCreated")

	con1, close1 := localPb.Subscribe(context.Background(), chn)
	con2, close2 := localPb.Subscribe(context.Background(), chn)

	localPb.Publish(context.Background(), chn, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), chn, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Con1:", (<-con1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Con2:", (<-con2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 7)
	close1()
	close2()
}

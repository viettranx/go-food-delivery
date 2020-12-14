package consumers

import (
	"context"
	"fooddlv/common"
	"fooddlv/common/asyncjob"
	"fooddlv/pubsub"
	"log"
)

type jobConsumer struct {
	evt *pubsub.Message
	hdl asyncjob.JobHandler
}

func RunSendNotificationAfterCreateNote(appCtx common.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.ChanNoteCreated)

	go func() {
		for {
			msg := <-c
			log.Println(msg.Data())
		}
	}()
}

func SendNotificationAfterCreateNote(appCtx common.AppContext, ctx context.Context) asyncjob.JobHandler {
	return func(ctx context.Context) error {
		//log.Println("Send notification for note:", msg.Data())
		return nil
	}
}

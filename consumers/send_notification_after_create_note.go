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

// Pubsub ---- Wrapper (Engine) ----- Group has Jobs (contains Consumer)

func SendNotificationAfterCreateNote(appCtx common.AppContext) consumerJob {
	return consumerJob{
		Title: "Send notification after create note",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			//appCtx.RealtimeEngine().BroadCastTo([]int{1, 2, 3}, "NoteCreated", "aaa")
			log.Println("SendNotificationAfterCreateNote ", message.Data())
			// Find all socket of userId and emit

			return nil
		},
	}
}

func SendEmailAfterCreateNote(appCtx common.AppContext) consumerJob {
	return consumerJob{
		Title: "Send email after create note",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			log.Println("SendEmailAfterCreateNote ", message.Data())
			return nil
		},
	}
}

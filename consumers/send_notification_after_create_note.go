package consumers

import (
	"context"
	"fooddlv/common"
	"log"
)

func RunSendNotificationAfterCreateNote(appCtx common.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.ChanNoteCreated)

	go func() {
		for {
			msg := <-c
			log.Println(msg.Data())
		}
	}()
}

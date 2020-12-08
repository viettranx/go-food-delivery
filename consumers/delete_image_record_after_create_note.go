package consumers

import (
	"context"
	"fooddlv/common"
	"log"
)

type HasImageIds interface {
	GetImageIds() []int
}

func RunDeleteImageRecordAfterCreateNote(appCtx common.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.ChanNoteCreated)

	go func() {
		for {
			msg := <-c
			if data, ok := msg.Data().(HasImageIds); ok {
				log.Println(data.GetImageIds())
			}
		}
	}()
}

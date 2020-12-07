package consumers

import (
	"context"
	"fooddlv/common"
)

func Setup(ctx common.AppContext) {
	RunDeleteImageRecordAfterCreateNote(ctx, context.Background())
	RunSendNotificationAfterCreateNote(ctx, context.Background())
}

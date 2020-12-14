package consumers

import (
	"context"
	"fooddlv/common"
	"fooddlv/pubsub"
)

func Setup(ctx common.AppContext) {
	RunDeleteImageRecordAfterCreateNote(ctx, context.Background())
	RunSendNotificationAfterCreateNote(ctx, context.Background())
}

type consumerEngine struct {
	appCtx common.AppContext
}

func NewEngine(appContext common.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appContext}
}

func (engine *consumerEngine) Start() error {
	//ps := engine.appCtx.GetPubsub()

	//engine.startSubTopic(common.ChanNoteCreated, asyncjob.NewGroup(
	//	false,
	//	asyncjob.NewJob(SendNotificationAfterCreateNote(engine.appCtx, context.Background(), nil))),
	//)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Channel, group GroupJob) error {

	//c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	//go func() {
	//	for {
	//		msg := <-c
	//		group.Run()
	//	}
	//}()

	return nil
}

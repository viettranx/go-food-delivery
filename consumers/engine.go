package consumers

import (
	"context"
	"fooddlv/common"
	"fooddlv/common/asyncjob"
	"fooddlv/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

func Setup(ctx common.AppContext) {

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

	engine.startSubTopic(
		common.ChanNoteCreated,
		false,
		SendNotificationAfterCreateNote(engine.appCtx),
		SendEmailAfterCreateNote(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Channel, isParallel bool, hdls ...consumerJob) error {
	// forever: listen message and execute group
	// hdls => []job
	// new group ([]job)

	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range hdls {
		log.Println("Setup consumer for:", item.Title)
	}

	getHld := func(job *consumerJob, message *pubsub.Message) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(hdls))

			for i := range hdls {

				// capture msg & hlds[i]
				//jobHdl := func(ctx context.Context) error {
				//	log.Println("running job for ", hdls[i].Title, ". Value: ", msg.Data())
				//	return hdls[i].Hld(ctx, msg)
				//}

				jobHdlArr[i] = asyncjob.NewJob(getHld(&hdls[i], msg))
			}

			group := asyncjob.NewGroup(isParallel, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}

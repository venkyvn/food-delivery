package subscriber

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component"
	"go-food-delivery/component/asyncjob"
	"go-food-delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx component.AppContext
}

func NewEngine(appContext component.AppContext) *consumerEngine {
	return &consumerEngine{
		appCtx: appContext,
	}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserUnLikeRestaurant,
		true,
		RunDecreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for", job.Title, ". Value :", message.Data())
			return job.Handler(ctx, message)
		}

	}

	go func() {
		for {
			msg := <-c

			jobHandlerArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHandler := getJobHandler(&consumerJobs[i], msg)
				jobHandlerArr[i] = asyncjob.NewJob(jobHandler)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHandlerArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}

		}
	}()

	return nil
}

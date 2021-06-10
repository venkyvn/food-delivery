package main

import (
	"context"
	"errors"
	"go-food-delivery/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println(" i am job 1")
		return errors.New("job 1 errrrrrrrr-----")
		//return nil
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 1)
		log.Println("I am job 2")

		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 1)
		log.Println("I am job 3")

		return nil
	})

	job4 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 1)
		log.Println("I am job 4")

		//return nil
		return errors.New("something went wrong at job 4")
	})

	group := asyncjob.NewGroup(true, job1, job2, job3, job4)
	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}

	//if err := job1.Execute(context.Background()); err != nil {
	//	log.Println("job 1 err: ", err)
	//
	//	for {
	//		if err := job1.Retry(context.Background()); err == nil {
	//			break
	//		}
	//
	//		log.Println("job 1 still error")
	//
	//		if job1.State() == asyncjob.StateRetryFailed {
	//			break
	//		}
	//	}
	//}
}

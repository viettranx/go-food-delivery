package main

import (
	"context"
	"errors"
	"fooddlv/common/asyncjob"
	"log"
	"time"
)

func main() {
	j1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("I am job 1")
		time.Sleep(time.Second * 5)
		return nil
	})

	j2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("I am job 2")
		time.Sleep(time.Second)
		return errors.New("err of job 2")
	})

	j2.SetRetryDurations([]time.Duration{time.Second * 2})

	group := asyncjob.NewGroup(true, j1, j2)
	err := group.Run(context.Background())

	log.Println("Group result:", err)
}

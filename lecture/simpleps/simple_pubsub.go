package main

import (
	"log"
	"time"
)

func startSender(name string, step int, queue chan int) {
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			i += step
			log.Printf("Sender %s sent: %d\n", name, i)
			queue <- i
		}
	}()
}

func startConsumer(name string, queue <-chan int) {
	go func() {
		for {
			time.Sleep(time.Second)
			log.Printf("Consumer %s received: %d\n", name, <-queue)
		}
	}()
}

func main() {
	queue := make(chan int, 100)

	startSender("s1", 1, queue)
	startSender("s2", 2, queue)

	startConsumer("c1", queue)
	startConsumer("c2", queue)

	time.Sleep(time.Second * 10)
}

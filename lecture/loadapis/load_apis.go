package main

import (
	"fmt"
	"time"
)

func fetchAPI(model string, duration time.Duration) string {
	time.Sleep(duration)
	return model
}

func main() {
	responseChan := make(chan string)
	var results []string

	go func() { responseChan <- fetchAPI("users", time.Second*4) }()
	go func() { responseChan <- fetchAPI("categories", time.Second*6) }()
	go func() { responseChan <- fetchAPI("products", time.Second*2) }()

	for i := 1; i <= 3; i++ {
		results = append(results, <-responseChan)
	}
	fmt.Println(results)
}

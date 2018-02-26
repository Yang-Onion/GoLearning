package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Printf("working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelAsync() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

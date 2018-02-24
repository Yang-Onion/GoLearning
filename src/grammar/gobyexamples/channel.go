package main

import "fmt"

func funcChannel() {
	message := make(chan string)
	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}

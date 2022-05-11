package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("Channel Work")
	case <-timeout:
		fmt.Println("Channel Timeout")
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("Channel Work")
	case <-time.After(1 * time.Second):
		fmt.Println("Channel Timeout")
	}
}

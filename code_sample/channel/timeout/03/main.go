package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 2
	select {
	case ch <- 2:
		fmt.Println("Channel Value is", <-ch)
	default:
		fmt.Println("channel blocked")
	}
}

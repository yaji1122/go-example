package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	//send
	go send(c)
	//receive
	receive(c)
}

//send something
func send(c chan<- string) {
	c <- "I'm a message"
}

func receive(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

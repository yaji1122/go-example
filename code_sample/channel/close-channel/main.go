package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//must close for range stop, or it will go into deadlock
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finish")
}

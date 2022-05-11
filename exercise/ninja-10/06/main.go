package main

import "fmt"

func main() {
	c := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Print(v, ",")
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for j := 0; j < 10; j++ {
					c <- j*10 + m
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()
	for v := range c {
		fmt.Print(v, ",")
	}
}

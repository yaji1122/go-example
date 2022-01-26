package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(100)

	counter := 0
	for i := 0; i < 100; i++ {
		go func() {
			runtime.Gosched()
			v := counter
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}

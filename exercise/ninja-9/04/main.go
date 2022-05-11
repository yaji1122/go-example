package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(100)

	counter := 0
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			runtime.Gosched()
			v := counter
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter :=  0

	const gs = 100
	//waitGroup
	var wg sync.WaitGroup
	//Mutex Lock
	var mutex sync.Mutex
	//How many goroutine to finish
	wg.Add(gs)

	for i:= 0 ; i < gs ; i++ {
		go func() {
			mutex.Lock()
			v := counter
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Printf("Goroutines(%d):%d \n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)


}

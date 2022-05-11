package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	foo()
	go bar()
	wg.Wait()
}

func foo() {
	fmt.Println("I'm a fool")
}

func bar() {
	fmt.Println("I'm bar")
	wg.Done()
}

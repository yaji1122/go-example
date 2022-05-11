package main

import "fmt"

var x interface{}

func main() {
	x = 1
	fmt.Printf("X is type of %T\n", x)
	x = "123"
	fmt.Printf("X is type of %T", x)
}

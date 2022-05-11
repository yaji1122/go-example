package main

import "fmt"

type gender int

var x gender

func main() {
	fmt.Println("value of x is", x)
	fmt.Printf("type of x is %T\n", x)
	x = 42
	fmt.Println("Again, value of x is", x)
}

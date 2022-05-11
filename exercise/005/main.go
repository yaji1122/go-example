package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println("value of x is", x)
	fmt.Printf("type of x is %T\n", x)
	x = 1
	fmt.Println("value of x is", x)
	y = int(x)
	fmt.Println("value of y is", y)
	fmt.Printf("type of y is %T", y)
}

package main

import "fmt"

func main() {
	defer finish()
	fmt.Println("Program start")
}

func finish() {
	fmt.Println("Exit Program")
}

package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println("Result is", a, b, c)
}

func foo() int {
	return 666
}

func bar() (int, string) {
	return 666, "ok"
}

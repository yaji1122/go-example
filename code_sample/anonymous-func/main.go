package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous func")
	}()

	func(x int) {
		fmt.Println("Anonymous func with parameter", x)
	}(42)

	f := func() {
		fmt.Println("function expression")
	}
	f()

	j := func(s string) {
		fmt.Println("function expression", s)
	}
	j("yup")

	h := foo()
	x := h()
	fmt.Println("func return a func return a int is", x)
}

func foo() func() int {
	return func() int {
		return 420
	}
}

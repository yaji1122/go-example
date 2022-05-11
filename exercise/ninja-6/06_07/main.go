package main

import "fmt"

func main() {
	//build an anonymous func and use it
	func() {
		fmt.Println("I'm anonymous func")
	}()
	//assign a func to a variable and calls it.
	f := func() {
		fmt.Println("I'm a func with a identifier")
	}
	f()
	//Create a func which returns a func
	//assign the returned func to a variable
	//call the returned func
	f2 := getFunc()
	f2()
	//pass a func into a func as an argument
	funcTrigger(f2)
}

func funcTrigger(f func()) {
	fmt.Println("Let me trigger a func")
	f()
}

func getFunc() func() {
	return func() {
		fmt.Println("I'm a nested func")
	}
}

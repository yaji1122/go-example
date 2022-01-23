package main

import "fmt"

func main() {
	//build an anonymous func and use it
	func(){
		fmt.Println("I'm anonymous func")
	}()
}

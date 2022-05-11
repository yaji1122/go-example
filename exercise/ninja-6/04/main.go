package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello mother fucker. My name is %s %s, %d years old. \n", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Yaji",
		last:  "Lin",
		age:   23,
	}
	p.speak()
}

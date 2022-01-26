package main

import "fmt"

type person struct {
	Name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("%s say: Hello. \n", p.Name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person {"Yaji"}
	saySomething(&p)
}



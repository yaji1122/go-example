package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type secretAgent struct {
	person
	weapon string
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("My name is", s.first_name, s.last_name)
}

func shit(h human) {
	fmt.Println("Oh I'm shitting.")
}

func main() {
	person1 := person{
		first_name: "James",
		last_name:  "Bond",
		age:        40,
	}

	agentBond := secretAgent{
		person: person1,
		weapon: "M41",
	}

	agentBond.speak()
	shit(agentBond)
	fmt.Printf("Type of agentBond is %T \n", agentBond)
}

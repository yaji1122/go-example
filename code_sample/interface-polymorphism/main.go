package main

import "fmt"

type human interface {
	speak()
}

func killed(h human) {
	switch h.(type) {
	case police:
		fmt.Printf("Citizen %s %s was killed. \n", h.(police).first, h.(police).last)
	case secretAgent:
		fmt.Printf("Secret Agent %s %s was killed. \n", h.(secretAgent).first, h.(secretAgent).last)
	}
}

type citizen struct {
	first string
	last  string
}

type secretAgent struct {
	citizen
	ltk    bool
	weapon string
}

type police struct {
	citizen
	weapon string
}

func (s secretAgent) speak() {
	fmt.Println("I'm secret agent", s.first, s.last)
}

func (p police) speak() {
	fmt.Println("I'm officer", p.first, p.last)
}

func main() {
	agent := secretAgent{
		citizen{
			first: "James",
			last:  "Bond",
		},
		true,
		"M41",
	}
	agent.speak()
	killed(agent)

}

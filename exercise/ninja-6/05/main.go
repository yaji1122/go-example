package main

import (
	"fmt"
	"math"
)

type square struct {
	l float32
	w float32
}

type circle struct {
	r float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.l * s.w
}

func (c circle) area() float32 {
	return c.r * c.r * math.Pi
}

func info(s shape) {
	fmt.Printf("The area of %T is %.2f \n", s, s.area())
}

func main() {
	s := square{
		l: 5.5,
		w: 4.5,
	}

	c := circle{
		r: 4.3,
	}

	info(s)
	info(c)
}

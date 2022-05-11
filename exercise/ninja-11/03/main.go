package main

import "fmt"

type customErr struct {
	s string
}

func (err customErr) Error() string {
	return err.s
}

func main() {
	e := customErr{"Oh no"}
	foo(e)
}

func foo(err error) {
	fmt.Println(err)
}

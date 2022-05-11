package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice is", foo(numbers...))
	fmt.Println("Sum of slice is", bar(numbers))
}

func foo(x ...int) int {
	sum := 0
	for v := range x {
		sum += v
	}
	return sum
}

func bar(numbers []int) int {
	sum := 0
	for x := range numbers {
		sum += x
	}
	return sum
}

package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	//for range
	for index, value := range x {
		fmt.Printf("Index[%d] is %d\n", index, value)
	}
	//slice an slice
	fmt.Println("Slice the slice")
	fmt.Println(x[1:3])
	//append element
	x = append(x, 84)
	fmt.Println(x)
	y := []int{123, 456, 789, 4234, 11573}
	x = append(x, y...)
	fmt.Println(x)
	//remove an element, remove the 3rd, 4th element
	x = append(x[:2], x[4:]...)
	getInfo(x)
	//make a slice with default length and capacity
	z := make([]int, 10, 100)
	getInfo(z)
	//Multi-dimentional slice
	groupOne := []string{"X", "Y", "Z"}
	groupTwo := []string{"A", "B", "C"}
	groupComplex := [][]string{groupOne, groupTwo}
	fmt.Println(groupComplex)
}

func getInfo(slice []int) {
	fmt.Println("Slice -> ", slice)
	fmt.Println("Length of slice : ", len(slice))
	fmt.Println("Capacity of slice : ", cap(slice))
}

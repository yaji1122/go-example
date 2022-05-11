package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort int
	i := []int{4, 5, 1, 7, 23, 0, -5}
	sort.Ints(i)
	fmt.Println(i)
	//sort string
	s := []string{"A", "D", "F", "B", "Banana", "Apple", "juice"}
	sort.Strings(s)
	fmt.Println(s)
	//find string
	index := sort.SearchStrings(s, "Banana")
	fmt.Println("index of banana is", index)
	//sort float
	f := []float64{512.33234, 0.233, 2.34, 4.028, 4123.123, 4512.42}
	sort.Float64s(f)
	fmt.Println(f)
}

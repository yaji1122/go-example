package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Last  string
	Age   int
}

type ByAge []user

type ByName []user

func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (b ByName) Len() int      { return len(b) }
func (b ByName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool {
	return b[i].First < b[j].First
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	// your code goes here
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByName(users))
	fmt.Println(users)

}

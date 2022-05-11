package main

import "fmt"

func main() {
	m := map[string]int{
		"Yaji":   33,
		"Jessie": 29,
	}
	// get a value by key
	age := m["Yaji"]
	fmt.Println(age)
	// if you get a value by a key that is not exist, it will return zero value of that type
	age = m["Null"]
	fmt.Println(age)
	//ask the key or value is exist in the map
	if _, exist := m["Null"]; !exist {
		fmt.Println("Null is not exist")
	}

	//range through a map
	for key, value := range m {
		fmt.Printf("%s is %d years old.\n", key, value)
	}

	//add value into the map
	m["John"] = 40
	fmt.Println(m)

	//remove an entry from map
	delete(m, "Yaji")
	fmt.Println(m)
}

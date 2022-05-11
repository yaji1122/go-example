package main

import "fmt"

//Exercise Part III
func main() {
	//#1 Print every number from 1 to 10000
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Println("")
	//#2 Print every rune code point of the uppercase alphabet three times.
	for i := 65; i < 91; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t %#U \n", i)
		}
	}
	//#3 Create a for loop using this syntax "for condition {}"
	init := 0
	for init < 10 {
		init++
	}
	fmt.Println("Finish Loop")
	//#4 Create a for loop using this syntax "for {}"
	init = 0
	for {
		init++
		if init > 10 {
			break
		}
	}
	fmt.Println("Finish Loop")
	//#5 Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println("")
	//#6 Create a program that shows the “if statement” in action.
	if init < 0 {
		fmt.Println("Init value is", init)
	}
	//#9
}

package main

import "fmt"

func main() {
	asciiNum := 36
	endNum := 122

	for asciiNum <= endNum {
		character := string(asciiNum)
		fmt.Printf("Number %d is ASCII [%s]\n", asciiNum, character)
		asciiNum++
	}
}

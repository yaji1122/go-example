package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `12345678`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(string(bs))

	pw2 := `3345678`
	err = bcrypt.CompareHashAndPassword(bs, []byte(pw2))
	if err != nil {
		fmt.Println("Error Password")
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//normal for loop
	var now time.Time = time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("time : ", time.Since(now))
	now = time.Now()

	//with index
	for i, arg := range os.Args {
		fmt.Printf("index of %d - %s \n", i, arg)
	}
	fmt.Println("time : ", time.Since(now))

	//join
	now = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println("time : ", time.Since(now))

}

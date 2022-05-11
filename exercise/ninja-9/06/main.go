package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s, Arch: %s \n", runtime.GOOS, runtime.GOARCH)
}

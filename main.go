package main

import (
	"fmt"

	"github.com/koheiterajima-bs/go-modules-hands-on/mymodule"
)

func main() {
	a := "Hello, Go Modules!"
	b := "Hello, Go Modules!"

	if mymodule.CompareStrings(a, b) {
		fmt.Println("Strings are equal!")
	} else {
		fmt.Println("Strings are not equal!")
	}
}

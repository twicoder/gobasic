package main

import (
	"fmt"
	"github.com/twicoder/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("\nHello, world2.\n"))
	a := "Hello"
	b := "World"
	c := a + b
	fmt.Printf(c)
}
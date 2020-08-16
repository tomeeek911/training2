package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

func foo() string {
	y := "Hello"
	return y
}

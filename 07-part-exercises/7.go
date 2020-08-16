package main

import "fmt"

func main() {
	a, b := foo()
	fmt.Println(a, b)
}

func foo() (string, int) {
	x := "Hey"
	z := 11
	return x, z
}

package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a())
	fmt.Printf("%T\n", a)
}

func foo() func() int {
	return func() int {
		return 21
	}
}

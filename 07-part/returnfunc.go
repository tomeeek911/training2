package main

import "fmt"

func main() {
	x := foo()
	fmt.Printf("%T\n", x)
	y := x()
	fmt.Println(y)
}

func foo() func() int {
	return func() int {
		return 555
	}
}

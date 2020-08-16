package main

import "fmt"

func main() {
	a := incrementor()
	fmt.Println(a())
	b := incrementor()
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x = 23
		return x
	}
}

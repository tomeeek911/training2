package main

import "fmt"

func main() {
	a := one()
	fmt.Printf("%T\n", a)
	fmt.Println(a())
}

func one() func() int {
	return func() int {
		var x int
		x = 21
		return x
	}

}

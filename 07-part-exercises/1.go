package main

import "fmt"

func main() {
	a := foo(42)
	fmt.Println(a)
	b, c := bar(24, "Tomasz")
	fmt.Println(b, c)

}

func foo(x int) int {
	return x
}

func bar(z int, u string) (int, string) {
	return z, u
}

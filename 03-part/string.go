package main

import "fmt"

func main() {
	a := "hello, this is Tomasz"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	ab := []byte(a)
	fmt.Println(ab)
	fmt.Printf("%T\n", ab)
}

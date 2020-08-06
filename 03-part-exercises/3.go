package main

import "fmt"

const (
	a int = 42           ///typed
	b     = 42.78        ///untyped
	c     = "James Bond" ///untyped
)

func main() {
	fmt.Println(a, b, c)
}

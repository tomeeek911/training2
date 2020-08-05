package main

import "fmt"

var z string = "This is a string"
var y string = `Tom said "hey this is actually a quote"`

func main() {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	z = "this too"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

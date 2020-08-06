package main

import "fmt"

var x int
var y float64

func main() {
	x = 1
	y = 1.23
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

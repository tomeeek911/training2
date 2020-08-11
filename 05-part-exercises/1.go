package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 52
	x[0] = 1
	x[4] = 13
	fmt.Println(x)
}

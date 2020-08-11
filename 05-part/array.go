package main

import (
	"fmt"
)

var x [10]int

func main() {
	x[3] = 42
	x[7] = 2
	fmt.Println(x)
}

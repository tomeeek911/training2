package main

import "fmt"

func main() {
	x := func(x int) {
		fmt.Println("This is my expression with an int equal to: ", x)
	}
	x(42)
}

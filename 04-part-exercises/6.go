package main

import "fmt"

func main() {
	x := 10
	if x == 3 {
		fmt.Println("10 is equal 3")
	}
	if x != 10 {
		fmt.Println("10 is not 10")
	}
	if x == 10 {
		fmt.Println("10 is equal 10 - true")
	}
}

package main

import "fmt"

func main() {
	x := 42
	if x == 41 {
		fmt.Println("wrong")
	} else if x == 42 {
		fmt.Println("correct")
	} else {
		fmt.Println("wrong")
	}
}

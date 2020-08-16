package main

import "fmt"

func main() {
	defer bar()
	foo()
	raw()
}
func foo() {
	fmt.Println("I am foo func")
}

func bar() {
	fmt.Println("I am defer func")
}

func raw() {
	fmt.Println("I am raw")
}

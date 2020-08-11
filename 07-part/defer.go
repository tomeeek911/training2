package main

import "fmt"

func main() {
	defer foo()
	bar()
	boom()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func boom() {
	fmt.Println("boom")
}

package main

import "fmt"

func main() {
	foo()
	bar("James")
	ok(42)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func ok(x int) {
	fmt.Println(x)
}

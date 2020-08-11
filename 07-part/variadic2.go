package main

import "fmt"

func main() {
	children("James", "Tomasz", "John", "Meggi", "Peter", "Bart")
}

func children(x ...string) {
	fmt.Println("Students with good grades:", x[0:2])
	fmt.Println("Students with bad grades:", x[2:])
	fmt.Println(x)
}

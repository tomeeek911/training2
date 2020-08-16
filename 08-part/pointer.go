package main

import "fmt"

func main() {
	i := 42
	fmt.Println("Before:", i)
	fmt.Println("Before:", &i)
	foo(&i)
	fmt.Println("After:", i)
	fmt.Println("After:", &i)
}

func foo(y *int) {
	fmt.Println("Before:", *y)
	fmt.Println("Before:", y)
	*y = 43
	fmt.Println("After:", *y)
	fmt.Println("After:", y)
}

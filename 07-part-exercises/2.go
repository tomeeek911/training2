package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4}
	a := foo(ii...)
	fmt.Println(a)
	ii2 := []int{1, 2, 3, 4, 5}
	b := bar(ii2)
	fmt.Println(b)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}

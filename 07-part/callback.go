package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := sum(a...)
	fmt.Println(b)
	b2 := even(sum, a...)
	fmt.Println(b2)
}

func sum(x ...int) int {
	summary := 0
	for _, v := range x {
		summary += v
	}
	return summary
}

func even(y func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return y(yi...)
}

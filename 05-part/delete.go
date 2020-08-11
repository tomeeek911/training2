package main

import "fmt"

func main() {
	x := []int{3, 4, 5, 6, 7, 8, 9}
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}

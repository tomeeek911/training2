package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)
}

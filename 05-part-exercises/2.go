package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}

}

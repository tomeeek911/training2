package main

import "fmt"

func main() {
	x := []int{2, 4, 76, 14}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])

	for i, v := range x {
		fmt.Println(i, v)
	}
}

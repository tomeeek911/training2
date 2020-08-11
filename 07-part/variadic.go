package main

import "fmt"

func main() {
	boom(2, 3, 4, 5)
}

func boom(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("This is the result from i:", i, "This is the result from v:", v, "Result from adding value to value:", sum)
	}
}

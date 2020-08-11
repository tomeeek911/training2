package main

import "fmt"

func main() {
	fmt.Println("Numbers devided by 4 between 10 and 100:")
	for x := 10; x <= 100; x++ {
		if x%4 == 0 {
			fmt.Printf("%d\n", x)
		}
	}
}

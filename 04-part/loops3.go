package main

import "fmt"

func main() {
	for x := 32; x <= 122; x++ {
		fmt.Printf("The decimal result: %d\t The characters result: %c\t The Unicode: %#U\n", x, x, x)
	}
}

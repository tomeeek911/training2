package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 1; y <= 3; y++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}

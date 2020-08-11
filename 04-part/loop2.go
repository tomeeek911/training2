package main

import "fmt"

func main() {
	for x := 0; x <= 20; x++ {
		if x == 19 {
			break
		}
		fmt.Println(x)
	}
}

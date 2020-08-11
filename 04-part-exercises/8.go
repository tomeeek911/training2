package main

import "fmt"

func main() {
	switch {
	case (1 == 1):
		fmt.Println("true")
	case (1 != 1):
		fmt.Println("false")
	}
}

package main

import "fmt"

func main() {
	children()
}

func children() {
	x := map[string]int{
		"Tomasz": 5,
		"Magda":  6,
		"John":   3,
		"Peter":  2,
	}
	for k, v := range x {
		if v < 4 {
			fmt.Println("This student:", k, ", has grades worse than 4:", v)
		}

	}
	for k, v := range x {
		if v >= 4 {
			fmt.Println("This student:", k, " has grades better than 4:", v)
		}
	}
}

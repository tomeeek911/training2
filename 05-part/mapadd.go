package main

import "fmt"

func main() {
	x := map[string]int{
		"James": 32,
		"MM":    27,
	}
	fmt.Println(x)
	x["Tomasz"] = 29
	for k, v := range x {
		fmt.Println(k, v)
	}
}

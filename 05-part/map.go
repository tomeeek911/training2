package main

import "fmt"

func main() {
	x := map[string]int{
		"James Bond": 32,
		"MM":         27,
	}
	fmt.Println(x)
	fmt.Println(x["James Bond"])
	fmt.Println(x["MM"])
}

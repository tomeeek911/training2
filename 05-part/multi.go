package main

import "fmt"

func main() {
	x := []string{"raz", "dwa", "trzy"}
	y := []string{"cztery", "piec", "szesc"}
	fmt.Println(x)
	fmt.Println(y)
	xy := [][]string{x, y}
	fmt.Println(xy)
}

package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "MoneyP", "Helo James"}
	xy := [][]string{x, y}
	fmt.Println(xy)
}

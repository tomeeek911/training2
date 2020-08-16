package main

import "fmt"

func main() {
	foo()
	func(x string) {
		fmt.Println(x)
	}("I am an anonymous function :lenny:")
	func() {
		x := map[string]int{
			"Tomasz": 29,
			"Magda":  28,
		}
		fmt.Println(x)
	}()
}
func foo() {
	x := "I am a foo function"
	fmt.Println(x)
}

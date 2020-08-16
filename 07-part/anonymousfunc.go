package main

import "fmt"

func main() {
	one()
	func() {
		fmt.Println(`This is an "anonymous func"`)
	}()
	func(x int) {
		fmt.Println(`This is an "anonymous func2", and its value is: `, x)
	}(42)

}

func one() {
	fmt.Println(`This is a "func one"`)
}

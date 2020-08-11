package main

import "fmt"

func main() {
	one()
	two("James")
	three()
	z := four("MM")
	fmt.Println(z)
	a, b := five("Tomasz, and he is: ", 27)
	fmt.Println(a)
	fmt.Println(b)
}

func one() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x[0:2])
	fmt.Println(x)
}

func two(s string) {
	fmt.Println("Hello,", s)
}

func three() {
	y := map[string]int{
		"James ma tyle lat:": 42,
		"John ma tyle lat":   27,
	}
	for k, v := range y {
		fmt.Println(k, v)
	}
}

func four(f string) string {
	return fmt.Sprint("hello, ", f, "!")
}

func five(g1 string, g2 int) (string, bool) {
	x := fmt.Sprint("This is, ", g1, ` `, g2)
	y := true
	return x, y
}

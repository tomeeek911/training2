package main

import "fmt"

func main() {

	x := struct {
		first string
		last  string
	}{
		first: "Tomasz",
		last:  "Sabiniewicz",
	}
	fmt.Println(x)
}

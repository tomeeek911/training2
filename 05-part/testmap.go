package main

import "fmt"

func main() {

	x := map[string]int{
		"Tomasz": 29,
		"Magda":  27,
	}
	fmt.Println("Imię i wiek:")
	for k, v := range x {
		fmt.Println(k, v)
	}
	y := map[string]string{
		"Tomasz": "samochód: Insignia",
		"Magda":  "samochód: Citroen",
	}
	fmt.Println(`
Imię i auto:`)
	for k, v := range y {
		fmt.Println(k, v)
	}
}

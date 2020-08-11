package main

import "fmt"

func main() {
	favSport := "hejka"
	switch {
	case favSport == "James Bond":
		fmt.Println("favSport is a Bond")
	case favSport != "James Bond":
		fmt.Println("favSport is not a Bond")
	case favSport == "hejka2":
		fmt.Println("favSport is hejka")
	}
}

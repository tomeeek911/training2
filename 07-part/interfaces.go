package main

import "fmt"

type geo struct {
	country    string
	city       string
	population int
}

type citizen struct {
	geo
	sex string
}

type human interface {
	stats()
}

func (a geo) stats() {
	fmt.Println("Here are information related to geolocalization: ", a.country, a.city, a.population)
}

func main() {
	x1 := geo{
		country:    "Poland",
		city:       "Poznan",
		population: 1000000,
	}
	x2 := citizen{
		geo: geo{
			country:    "England",
			city:       "London",
			population: 5000000,
		},
		sex: "male",
	}
	fmt.Println(x1)
	fmt.Println(x2)
	x1.stats()
	fmt.Printf("%T\n", x1)
}

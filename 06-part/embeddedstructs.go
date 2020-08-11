package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type woman struct {
	person
	wom bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   27,
	}
	p2 := person{
		first: "M",
		last:  "M",
		age:   21,
	}
	w1 := woman{
		person: person{
			first: "James",
			last:  "Bond",
			age:   27,
		},
		wom: false,
	}
	w2 := woman{
		person: person{
			first: "M",
			last:  "M",
			age:   21,
		},
		wom: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(w1)
	fmt.Println(w2)
	fmt.Println(w1.first, w1.last, w1.age, w1.wom)
	fmt.Println(w2.first, w2.last, w2.age, w2.wom)

}

package main

import "fmt"

type person struct {
	name     string
	sourname string
	age      int
	male     bool
}

type car struct {
	sport bool
	color string
}

type employee struct {
	person
	car
}

func main() {
	first(1, 2, 3, 4, 5)
	second()
	e1 := employee{
		person: person{
			name:     "Tomasz",
			sourname: "Sabiniewicz",
			age:      29,
			male:     true,
		},
		car: car{
			sport: true,
			color: "red",
		},
	}
	e2 := employee{
		person: person{
			name:     "Magda",
			sourname: "Sabiniewicz",
			age:      28,
			male:     false,
		},
		car: car{
			sport: false,
			color: "grey",
		},
	}
	e3 := employee{
		person: person{
			name:     "Jaś",
			sourname: "Sabiniewicz",
			age:      0,
			male:     true,
		},
		car: car{
			sport: false,
			color: "white",
		},
	}

	fmt.Println(`
1st employee: `)
	fmt.Println("Employee's name: ", e1.name, e1.sourname)
	fmt.Println("Employee's age: ", e1.age)
	fmt.Println("Male: ", e1.male)
	fmt.Println("Employee's car color: ", e1.color)
	fmt.Println("Sports car: ", e1.sport)
	fmt.Println(`
2nd employee: `)
	fmt.Println("Employee's name: ", e2.name, e2.sourname)
	fmt.Println("Employee's age: ", e2.age)
	fmt.Println("Male: ", e2.male)
	fmt.Println("Employee's car color: ", e2.color)
	fmt.Println("Sports car: ", e2.sport)
	e3.third()

	h := fourth("Kolesław Nowag")
	fmt.Println(h)

	fifth()

}

func first(x ...int) {
	fmt.Println("Variadic parameter: ", x)
	fmt.Printf("%T\n", x)
}

func second() {
	x := []string{
		"James",
		"Tom",
		"Magda",
	}
	fmt.Println("The first string under this slice is: ", x[0])
	fmt.Println("The second string under this slice is: ", x[1])
	fmt.Println("The third string under this slice is: ", x[2])

}

func (t employee) third() {
	fmt.Println(`
I am a third employee: `, t.name, t.sourname)
}

func fourth(h string) string {
	return fmt.Sprint(`
Hey, I am: `, h)
}

func fifth() {
	a := map[string]int{
		"James":  28,
		"John":   29,
		"Ann":    45,
		"Joanne": 32,
	}
	for k, v := range a {
		if v < 30 {
			fmt.Println(`Younger than 30: `, k, v)
		}
	}
	for k, v := range a {
		if v > 30 {
			fmt.Println("Older than 30: ", k, v)
		}
	}
}

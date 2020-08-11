package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheels: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)
	fmt.Println(t1.vehicle.doors, t1.vehicle.color, t1.fourWheels)

	fmt.Printf("%T\n", s1)
	fmt.Println(s1)
	fmt.Println(s1.vehicle.doors, s1.vehicle.color, s1.luxury)

}

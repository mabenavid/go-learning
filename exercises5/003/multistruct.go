package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	dTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "gray",
		},
		fourWheel: true,
	}

	dSedan := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		luxury: false,
	}

	fmt.Println(dTruck)
	fmt.Println(dTruck.doors, dTruck.color, dTruck.fourWheel)
	fmt.Println(dSedan)
	fmt.Println(dSedan.doors, dSedan.color, dSedan.luxury)

}

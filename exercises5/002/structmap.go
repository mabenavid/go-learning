package main

import "fmt"

type person struct {
	firstName, lastName string
	favoriteFlavors     []string
}

func main() {

	p1 := person{
		firstName:       "Manu",
		lastName:        "Bena",
		favoriteFlavors: []string{"vanilla", "chocolate"},
	}

	p2 := person{
		firstName:       "Yuli",
		lastName:        "Loria",
		favoriteFlavors: []string{"bubble gum", "choco menta"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		for j, val := range v.favoriteFlavors {
			fmt.Println(j, val)
		}
	}
}

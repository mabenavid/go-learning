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

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favoriteFlavors {
		fmt.Printf("\t%v\n", v)

	}
	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.favoriteFlavors {
		fmt.Printf("\t%v\n", v)
	}
}

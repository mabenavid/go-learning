package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Manu2"
	p.last = "Bena2"
	p.age = 39
	//OR
	(*p).first = "Manu3"
	(*p).last = "Bena3"
	(*p).age = 40
}

func main() {

	p1 := person{
		first: "Manu",
		last:  "Bena",
		age:   38,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)

}

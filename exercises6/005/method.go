package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.l * s.w
}

func (c circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		l: 3.1,
		w: 2.2,
	}
	c := circle{
		r: 2.33,
	}

	info(s)
	info(c)
}

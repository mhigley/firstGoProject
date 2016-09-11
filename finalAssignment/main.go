package main

import (
	"fmt"
	"math"
)

/*
type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println(p.fName, ` says... wuzzup`)
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fName, ` says... you never saw anything`, sa.licenseToKill)
}

func main() {
	p1 := person{
		"Matt",
		"Higley",
	}
	fmt.Println(p1.lName)

	p2 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(p2.licenseToKill)
}
*/

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println(x)
	fmt.Println(x.area())
}

func main() {
	s := square{20}
	info(s)

	c := circle{15}
	info(c)
}

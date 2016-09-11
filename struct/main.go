package main

import "fmt"

type Person struct {
	Fname   string
	Lname   string
	Age     int
	favNums []int
	fPh     map[string]int
}

func main() {
	//p1 := Person{"Todd", "McLeod"}
	p1 := Person{
		Fname:   "Matt",
		Lname:   "Higley",
		Age:     45,
		favNums: []int{7, 14, 57, 42, 451},
		fPh:     map[string]int{"Bob": 42, "Matt": 40},
	}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
}

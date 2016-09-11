package main

import "fmt"

func main() {
	var xs []string
	xs = append(xs, "Matthew")
	xs = append(xs, "Higley")
	fmt.Println(xs)

	xb := []bool{true, false, true} //composite literal syntax
	fmt.Println(xb)

	/*
		y := make([]int, 0, 2)
		fmt.Println(y)
		fmt.Printf("%T\n", y)
		fmt.Println("length", len(y))
		fmt.Println("capacity", cap(y))
		fmt.Println("~~~~~~~~~~~")

		y = append(y, 14)
		fmt.Println(y)
		fmt.Printf("%T\n", y)
		fmt.Println("length", len(y))
		fmt.Println("capacity", cap(y))
		fmt.Println("~~~~~~~~~~~")

		y = append(y, 17)
		fmt.Println(y)
		fmt.Printf("%T\n", y)
		fmt.Println("length", len(y))
		fmt.Println("capacity", cap(y))
		fmt.Println("~~~~~~~~~~~")

		y = append(y, 19)
		fmt.Println(y)
		fmt.Printf("%T\n", y)
		fmt.Println("length", len(y))
		fmt.Println("capacity", cap(y))
		fmt.Println("~~~~~~~~~~~")
	*/
}

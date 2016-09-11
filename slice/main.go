package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := 451
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var z []int
	z = append(z, x)
	z = append(z, y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(z[0])
	fmt.Println(z[1])

	nums := []int{7, 9, 37, 16}
	fmt.Print(nums)
}

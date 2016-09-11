package main //required

import "fmt" //imports

func init() { //initialize the setup for your program
	fmt.Println("this will run first")
}
func main() { //required
	fmt.Println("Hello World!")
}

// go run
// go build
// run with: ./`binary name`
// go install

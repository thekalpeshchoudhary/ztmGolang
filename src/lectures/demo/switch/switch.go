package main

import "fmt"

func price() int {
	return 3
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap")
	case p < 10:
		fmt.Println("Moderate")
	default:
		fmt.Println("Expensive")
	}

	ticket := FirstClass
	switch ticket {
	case Economy:
		fmt.Println("Economy Seats")
	case Business:
		fmt.Println("Buisness Seats")
	case FirstClass:
		fmt.Println("FirstClass Seats")
	}

}

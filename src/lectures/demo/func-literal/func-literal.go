package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, opp func(lhs, rhs int) int) int {
	fmt.Printf("Runnning a computation with %v and  %v \n", lhs, rhs)
	return opp(lhs, rhs)
}
func main() {
	fmt.Println("5 + 4 :",
		compute(5, 4, add))

	fmt.Println("10 - 2 :", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("multiplying %v with %v = ", lhs, rhs)
		return lhs * rhs
	}
	fmt.Println(compute(4, 9, mul))
}

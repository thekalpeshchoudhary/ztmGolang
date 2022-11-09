package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, eachNum := range nums {
		sum += eachNum
	}
	return sum
}

func main() {
	fmt.Println("Sum1", sum(1, 2, 4, 1))
	a := []int{1, 3}
	b := []int{1, 2}

	all := append(a, b...)
	fmt.Println(sum(all...))
}

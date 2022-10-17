package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 2, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("'Q1 is higher than Q2")
	} else if quiz2 > quiz1 {
		fmt.Println("Q2 is higher than Q1")
	} else {
		fmt.Println("Q1 and Q2 are same")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Failed")
	}
}

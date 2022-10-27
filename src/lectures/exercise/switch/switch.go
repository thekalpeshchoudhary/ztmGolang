//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {

	switch age := 19; {
	case age == 0:
		fmt.Println("newborn")
		//  - "newborn" when age is 0
	case 1 >= age && age <= 3:
		fmt.Println("toddler")
		//  - "toddler" when age is 1, 2, or 3
	case age >= 4 && age <= 12:
		fmt.Println("child")
		//  - "child" when age is 4 through 12
	case age >= 13 && age <= 17:
		fmt.Println("teenager")
		//  - "teenager" when age is 13 through 17
	default:
		fmt.Println("adult")
		//  - "adult" when age is 18+
	}

}

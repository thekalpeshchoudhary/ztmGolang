//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "Red"
	fmt.Println("Color is :", color)

	year, age := 1998, 24
	fmt.Println("Year ",year, "age",age)

	var (
		firstInitial = "K"
		lastInitial = "C"
	)
	fmt.Println(firstInitial,lastInitial)

	var myAge int
	myAge = age * 365
	fmt.Println("Age in days", myAge)

}

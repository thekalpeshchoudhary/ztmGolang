//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string){
	fmt.Println("Hello there",name,"!")
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func returnMessage() string {
	return "Random Message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addNumbeers(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

//* Write a function that returns any number
func anyNum() int {
	return 4
}

//* Write a function that returns any two numbers
func anyTwoNum() (int,int) {
	return 3, 4
}

func main() {
	greet("Kalpesh")
	
	randMsg := returnMessage()
	fmt.Println(randMsg)
	
	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	sumOfThree := addNumbeers(1,2,3)
	fmt.Println("3 number sum is",sumOfThree)

	num1, num2 := anyTwoNum()
	sumOfThreeAgain := anyNum()+ num1+ num2
	fmt.Println("Sum of 3 numbers is",sumOfThreeAgain)
}

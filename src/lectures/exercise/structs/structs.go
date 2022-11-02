//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func length(rectAngle Rectangle) int {
	return (rectAngle.a.y - rectAngle.b.y)
}
func width(rectAngle Rectangle) int {
	return (rectAngle.b.x - rectAngle.a.x)
}
func calcArea(rectAngle Rectangle) int {
	return length(rectAngle) * width(rectAngle)
}
func calcPerimeter(rectAngle Rectangle) int {
	return (length(rectAngle) * 2) + (width(rectAngle) * 2)
}

func printInfo(rectAngle Rectangle) {
	fmt.Println("Rectangle", rectAngle)
	fmt.Println("Area :", calcArea(rectAngle))
	fmt.Println("Perimeter :", calcPerimeter(rectAngle))
}

func main() {
	rectAngle := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rectAngle)

	rectAngle.a.y *= 2
	rectAngle.b.x *= 2
	printInfo(rectAngle)

}

//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import (
	"fmt"
)

// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Car string
type Truck string
type Bike string

// * Vehicles have a model name in addition to the vehicle type:
//   - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
func (b Bike) String() string {
	return fmt.Sprintf("Bike : %v", string(b))
}
func (c Car) String() string {
	return fmt.Sprintf("Car : %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck : %v", string(t))
}

func (b Bike) PickLift() Lift {
	return SmallLift
}
func (c Car) PickLift() Lift {
	return StandardLift
}
func (t Truck) PickLift() Lift {
	return LargeLift
}

//   - Write a single function to handle all of the vehicles
//     that the shop works on.
func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to Small lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to Standard lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to Large Lift lift\n", p)
	}
}

func main() {
	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	car := Car("Nexon")
	truck := Truck("Ashok Leyland")
	bike := Bike("KTM")

	sendToLift(bike)
	sendToLift(car)
	sendToLift(truck)
}

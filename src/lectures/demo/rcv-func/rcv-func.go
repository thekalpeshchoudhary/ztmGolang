package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println("Lot :", lot)

	lot.occupySpace(2)
	occupySpace(&lot, 1)
	fmt.Println("Lot :", lot)

	lot.vacateSpace(2)
	fmt.Println("Lot :", lot)

}

package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	kc := Passenger{"Kalpesh", 1, false}
	fmt.Println(kc)

	var (
		bill  = Passenger{Name: "Bill", TicketNumber: 2}
		hella = Passenger{Name: "Hella", TicketNumber: 3}
	)
	fmt.Println(bill, hella)

	var rohan Passenger
	rohan.Name = "Rohan"
	rohan.TicketNumber = 4
	fmt.Println(rohan)

	bill.Boarded = true
	kc.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill is Boarded")
	}
	if kc.Boarded {
		fmt.Println(kc.Name, "is Boarded")
	}

	hella.Boarded = true
	bus := Bus{hella}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the frontseat")
}

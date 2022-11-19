package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook Chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("Chop")
	fmt.Println("Add Dressing")
}
func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare Dishes :")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v \n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken Wings"), Salad("Iceberg Salad"),
	}
	prepareDishes(dishes)

}

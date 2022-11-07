package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["bread"] += 1
	shoppingList["milk"] = 1

	shoppingList["eggs"] += 1

	delete(shoppingList, "milk")
	fmt.Println("List with Milk removed:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal not found")
	} else {
		fmt.Println("Cereal:", cereal)
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("Total Items", totalItems)
}

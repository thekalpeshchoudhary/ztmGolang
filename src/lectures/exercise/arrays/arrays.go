//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//  - Products must include the price and the name
type Product struct {
	name  string
	price int
}

//* Print to the terminal:
func printShoppingList(shoppingList [4]Product) {
	var items, total int
	for i := 0; i < len(shoppingList); i++ {
		eachItem := shoppingList[i]
		total += eachItem.price
		if eachItem.name != "" {
			items += 1
		}
	}
	//  - The last item on the list
	fmt.Println("Last Item :", shoppingList[items-1])
	//  - The total number of items
	fmt.Println("Total Price", total)
	//  - The total cost of the items
	fmt.Println("Total number of Items:", items)
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products

	//* Insert 3 products into the array
	shoppingList := [4]Product{
		{name: "Phone", price: 50000},
		{name: "Charger", price: 3200},
		{name: "Headphones", price: 8850},
	}
	printShoppingList(shoppingList)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = Product{name: "Watch", price: 25000}
	printShoppingList(shoppingList)

}

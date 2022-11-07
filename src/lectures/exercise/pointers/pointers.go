//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//  - Security tags have two states: active (true) and inactive (false)
type SecurityTag bool

//* Create a structure to store items and their security tag state
type ItemTag struct {
	item string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func changeTagState(item *ItemTag) {
	item.tag = !item.tag
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []ItemTag) {
	for i := range items {
		if items[i].tag {
			changeTagState(&items[i])
		}
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	items := []ItemTag{
		{"phone", true},
		{"laptop", true},
		{"radio", true},
		{"radio", true},
	}

	//  - Deactivate any one security tag in the array/slice
	changeTagState(&items[2])
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println(items)

	//  - Print out the array/slice after each change
}

package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no element at index %v", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	values := []int{1, 2, 3}
	stuff := Stuff{values}
	value, err := stuff.Get(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}

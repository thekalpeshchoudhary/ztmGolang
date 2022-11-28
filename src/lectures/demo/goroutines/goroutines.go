package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitailzed []rune

	capIt := func(r rune) {
		capitailzed = append(capitailzed, unicode.ToUpper(r))
		fmt.Printf("%c Done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Capitalized %c", capitailzed)
}

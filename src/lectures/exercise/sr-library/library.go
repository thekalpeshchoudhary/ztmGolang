//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Member string
type Title string

type Book struct {
	inLibrary  bool
	checkedIn  string
	checkedOut string
}

type Books map[Title]Book

// The library must have books and members, and must include:
type Library struct {
	//  - Which books have been checked out
	books   Books
	members []Member
}

func libraryState(library Library) {
	fmt.Println("\nLibrary Stats")
	books := library.books
	for book, bookState := range books {
		fmt.Println("\nBook :", book)
		if bookState.inLibrary {
			fmt.Println("Book Checked In at :", bookState.checkedIn)
		} else {
			fmt.Println("Book Checked Out at :", bookState.checkedOut)
		}
	}

	members := library.members
	for _, member := range members {
		fmt.Println("Member Name :", member)
	}
}

func bookCheckout(title Title, library *Library) {
	book, found := library.books[title]
	if found {
		book.inLibrary = false
		book.checkedOut = time.Now().String()
		library.books[title] = book
		fmt.Println("Book Checked Out :", title)
	} else {
		fmt.Println("Book not found :", title)
	}
}

func bookCheckIn(title Title, library *Library) {
	book, found := library.books[title]
	if found {
		book.inLibrary = true
		book.checkedIn = time.Now().String()
		library.books[title] = book
		fmt.Println("Book Checked In :", title)
	} else {
		fmt.Println("Book not found :", title)
	}
}

func main() {
	//  - Add at least 4 books and at least 3 members to the library
	theBooks := Books{
		"alchemist":       Book{inLibrary: true, checkedIn: time.Now().String()},
		"do_epic_shit":    Book{inLibrary: true, checkedIn: time.Now().String()},
		"harry_potter":    Book{inLibrary: true, checkedIn: time.Now().String()},
		"game_of_thrones": Book{inLibrary: true, checkedIn: time.Now().String()},
	}

	theMembers := []Member{
		"Kalpesh", "Junaid", "Ishita",
	}

	library := Library{
		books:   theBooks,
		members: theMembers,
	}

	libraryState(library)

	//  - Check out a book
	bookCheckout("harry_potter1", &library)
	libraryState(library)

	//  - Check in a book
	bookCheckIn("alchemist", &library)
	libraryState(library)

	//  - Print out initial library information, and after each change
}

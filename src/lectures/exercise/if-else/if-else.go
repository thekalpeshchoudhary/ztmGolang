//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekday(today int) bool {
	return today < 5
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Monday, Member

	if role < 30 {
		//* Access at any time: Admin, Manager
		accessGranted()
	} else if role == 30 && !weekday(today) {
		//* Access weekends: Contractor
		accessGranted()
	} else if role == 40 && weekday(today) {
		//* Access weekdays: Member
		accessGranted()
	} else if role == 50 && (today == 0 || today == 2 || today == 4) {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
	} else {
		accessDenied()
	}
}

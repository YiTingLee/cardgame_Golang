package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	contactInfo //=== contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	eric := person{
		firstName: "Eric",
		lastName:  "Lee",
		contactInfo: contactInfo{
			email:   "eric@mail.com",
			zipCode: 200,
		},
	}

	// var eric person
	// eric.firstName = "Eric"
	// eric.lastName = "Lee"
	// eric.contact.email = "eric@mail.com"
	// eric.contact.zipCode = 2000

	// ericPointer := &eric
	// ericPointer.updateName("Gary")
	eric.updateName("Gary")
	eric.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

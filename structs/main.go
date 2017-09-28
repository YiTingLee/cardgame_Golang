package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	eric := person{
		firstName: "Eric",
		lastName:  "Lee",
		contact: contactInfo{
			email:   "eric@mail.com",
			zipCode: 200,
		},
	}

	// var eric person
	// eric.firstName = "Eric"
	// eric.lastName = "Lee"
	// eric.contact.email = "eric@mail.com"
	// eric.contact.zipCode = 2000

	fmt.Printf("%+v", eric)
}

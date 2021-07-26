package main

import "fmt"

type contact struct {
	email   string
	address string
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact
}

func main() {
	bilbo := person{
		firstName: "Bilbo",
		lastName:  "Baggins",
		age:       111,
		contact: contact{
			email:   "bilbo@gmail.com",
			address: "The shire",
		},
	}
	fmt.Printf("%+v", bilbo)
}

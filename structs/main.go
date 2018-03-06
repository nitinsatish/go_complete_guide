package main

import "fmt"

type contact struct {
	email   string
	zipcode int
}

type people struct {
	firstName   string
	lastName    string
	contactinfo contact
}

func main() {

	nitin := people{
		firstName:   "Nitin",
		lastName:    "Satish",
		contactinfo: contact{email: "nits@mail.com", zipcode: 126},
	}

	fmt.Printf("%+v\n", nitin)
	fmt.Println(nitin.contactinfo)
	(&nitin).updateFName("Nits")
	fmt.Printf("%v\n", nitin)

}

func (p *people) updateFName(newName string) {
	(*p).firstName = newName
}

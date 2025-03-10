package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "jim@anderson.com",
			zipCode: 12345,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()

    jim.printContact()

}

func (p person) printContact()  {
    fmt.Println("Email addr: ", p.email, "\nZip code:", p.zipCode)
}

// Struct receiver
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

// this also works
//
//	type person struct {
//		firstName string
//		lastName  string
//		contactinfo contactInfo
//	}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// name := person{firstName: "Roopesh", lastName: "Kp"}
	// fmt.Println(name)
	// var name person
	// name.firstName = "roopesh"
	// name.lastName = "kp"

	// fmt.Printf("%+v", name)
	name := person{firstName: "Roopesh", lastName: "Kp", contactInfo: contactInfo{email: "roopesh@gmail.com", zipcode: 89888}}
	// fmt.Printf("%+v", name)
	// namePointer := &name
	name.updatePerson("New name")
	name.print()
}
func (p *person) updatePerson(newName string) {
	(*p).firstName = newName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

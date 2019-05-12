package main

import "fmt"

type contactInfo struct {
	email  string
	mobile string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	mark := person{firstName: "Mark", lastName: "Ben"}
	var name person
	name.firstName = "Ben"
	name.lastName = "Afleck"
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:  "jim.party@gmail.com",
			mobile: "945384941",
		},
	}
	fmt.Println(mark, name, jim)
	jimpointer := &jim              //1
	jimpointer.updateName("jimmy1") //2
	//above two lines can also be  written as jim.updateName("jimmy")
	jim.print()
}
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

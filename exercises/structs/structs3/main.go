// structs3
// Make me compile!
//
// I AM DONE
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (person Person) FullName() string {
	return person.firstName + " " + person.lastName

}

func main() {
	person := Person{firstName: "Maurício", lastName: "Antunes"}
	fmt.Printf("Person full name is: %s\n", person.FullName()) // here it must output Person full name is: Maurício Antunes
}

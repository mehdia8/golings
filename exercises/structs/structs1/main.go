// structs1
// Make me compile!
//
// I AM DONE
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{name: "per", age: 20}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}

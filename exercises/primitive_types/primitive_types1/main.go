// primitive_types1
// Make me compile!

// I AM DONE
package main

import "fmt"

func main() {
	storeIsOpen := true
	if storeIsOpen {
		fmt.Println("The store is open, let's buy some clothes!")
	}

	storeIsOpen = false
	if !storeIsOpen {
		fmt.Println("Oh no, let's buy some clothes online!")
	}
}

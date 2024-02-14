// generics1
// Make me compile!

// I AM DONE
package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print[V string | int](value V) {
	fmt.Println(value)
}

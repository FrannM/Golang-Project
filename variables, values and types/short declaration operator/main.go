package main

import "fmt"

// DECLARE the VARIABLE "y"
// ASSIGN the value 43
// DECLARE & ASSIGN = initialization
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for numeric types, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}

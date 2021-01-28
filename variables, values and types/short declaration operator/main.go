package main

import "fmt"

var y = 43

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}

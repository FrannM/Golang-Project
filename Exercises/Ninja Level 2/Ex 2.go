package main

import "fmt"

func main() {
	x := 10
	y := 50
	a := (y == x)
	b := (x <= y)
	c := (x >= y)
	d := (x != y)
	e := (y < x)
	f := (y > x)

	fmt.Println(a, b, c, d, e, f)

}

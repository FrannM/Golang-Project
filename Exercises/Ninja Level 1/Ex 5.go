package main

import "fmt"

type num int

var b num
var c int

func main() {
	fmt.Println(b)
	fmt.Printf("%T\n", x)
	b = 42
	fmt.Println(b)
	c = int(b)
	fmt.Println(c)
	fmt.Printf("%T\n", y)

}

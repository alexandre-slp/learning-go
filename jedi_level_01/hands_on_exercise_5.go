package main

import (
	"fmt"
)

type alex int  //int is the underlying type

var x alex
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	y = int(x)
	fmt.Println(y)
}

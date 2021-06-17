package main

import "fmt"

type alex int  //int is the underlying type

var x alex

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

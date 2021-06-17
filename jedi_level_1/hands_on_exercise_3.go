package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	//s := fmt.Sprintf("%d, %s, %t", x, y, z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)
}

package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("decimal: %d\n" +
		"binary: %b\n" +
		"hex: %#x",
		x, x, x)
}

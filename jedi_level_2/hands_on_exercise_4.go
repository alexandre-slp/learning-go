package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("binary: %b\n", a)
	fmt.Printf("decimal: %d\n", a)
	fmt.Printf("hex: %#x\n", a)

	shifted := a << 1

	fmt.Printf("binary: %b\n", shifted)
	fmt.Printf("decimal: %d\n", shifted)
	fmt.Printf("hex: %#x\n", shifted)
}

package main

import "fmt"

func main() {
	f := foo()
	fmt.Printf("%T\n", f)
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
package main

import "fmt"

func main() {
	f := foo(2)
	g := foo(2)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
}

func foo(i int) func() int {
	vx := 0
	return func() int {
		vx += i
		return vx
	}
}

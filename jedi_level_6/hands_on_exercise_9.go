package main

import "fmt"

func main() {
	f := foo(func(x int) int {
		return 2 * x
	}, 2)
	fmt.Println(f)
}

func foo(f func(x int) int, x int) int {
	return f(x)
}

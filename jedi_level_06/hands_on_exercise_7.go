package main

import "fmt"

func main() {
	f := func(x int) int {
		return 2 * x
	}
	fmt.Printf("%T\n", f)
	fmt.Println(f(2))
}

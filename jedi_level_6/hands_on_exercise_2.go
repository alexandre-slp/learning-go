package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(foo(numbers...))
	fmt.Println(bar(numbers))
	fmt.Println("main ended")
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
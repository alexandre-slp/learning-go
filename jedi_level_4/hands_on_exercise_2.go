package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", s)
}

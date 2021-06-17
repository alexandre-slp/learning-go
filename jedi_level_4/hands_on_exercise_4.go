package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s1 := append(x, 52)
	fmt.Println(s1)

	s2 := append(x, 53, 54, 55)
	fmt.Println(s2)

	y := []int{56, 57, 58, 59, 60}
	s3 := append(x, y...)
	fmt.Println(s3)
}

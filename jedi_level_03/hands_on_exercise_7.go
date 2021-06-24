package main

import "fmt"

func main() {
	x := 21
	if x < 21 {
		fmt.Println("Age below 21")
	} else if x == 21 {
		fmt.Println("Age equal 21")
	} else {
		fmt.Println("Age over 21")
	}
}

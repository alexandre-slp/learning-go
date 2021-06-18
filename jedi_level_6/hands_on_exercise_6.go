package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println(x)
	}(2)
}

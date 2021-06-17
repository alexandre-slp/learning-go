package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not print")
	case true:
		fmt.Println("Print")
	}
}

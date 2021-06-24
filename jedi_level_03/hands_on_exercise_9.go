package main

import "fmt"

func main() {
	favoriteSport := "Soccer"
	switch favoriteSport {
	case "Basketball":
		fmt.Println("Not print")
	case "Soccer":
		fmt.Println("Print")
	default:
		fmt.Println("Not found")
	}
}

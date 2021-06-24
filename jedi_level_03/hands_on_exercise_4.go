package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	birthdate := 1990
	i := birthdate
	for {
		if i > now.Year() {
			break
		}
		fmt.Println(i)
		i++
	}
}

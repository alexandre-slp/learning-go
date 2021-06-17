package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	birthdate := 1990
	i := birthdate
	for i <= now.Year(){
		fmt.Println(i)
		i++
	}
}

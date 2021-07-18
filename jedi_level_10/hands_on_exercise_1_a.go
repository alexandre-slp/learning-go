package main

import (
	"fmt"
)

//var c = make(chan int)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

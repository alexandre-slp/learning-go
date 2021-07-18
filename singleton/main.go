package main

import (
	"fmt"
	"github.com/alexandre-slp/learning-go/singleton/clock"
)

func main() {
	c1 := clock.New()
	c2 := clock.New()

	fmt.Printf("%v\n", c1.Now())
	fmt.Println(c2)
}

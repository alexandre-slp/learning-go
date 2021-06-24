package main

import (
	"fmt"
)

func main() {
	numChan := make(chan int)
	stopChan := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			numChan <- i
		}
		stopChan <- 1
	}()

	for  {
		select {
		case v :=  <- numChan:
			fmt.Println(v)
		case <- stopChan:
			return
		}
	}
}

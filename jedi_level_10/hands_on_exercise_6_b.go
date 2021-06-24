package main

import "fmt"

func main() {
	numChan := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for v := range numChan {
		fmt.Println(v)
	}
}

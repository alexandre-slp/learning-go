package main

import "fmt"

func main() {
	numChan := make(chan int)
	totalGoroutines := 10

	for i := 0; i < totalGoroutines; i++ {
		go func(numGor int) {
			for j := 0; j < 10; j++ {
				numChan <- j
			}
			if numGor == totalGoroutines-1 {
				close(numChan)
			}
		}(i)
	}

	for v := range numChan {
		fmt.Println(v)
	}
}

package main

import "fmt"

func main() {
	numChan := make(chan int)

	for i := 0; i < 10; i++ {
		go func(numGor int) {
			for j := 0; j < 10; j++ {
				numChan <- j
			}
			if numGor == 9 {
				close(numChan)
			}
		}(i)
	}

	for v := range numChan {
		fmt.Println(v)
	}
}

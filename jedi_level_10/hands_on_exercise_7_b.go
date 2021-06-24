package main

import "fmt"

func main() {
	numChan := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				numChan <- j
			}
		}()
	}

	for k := 0;;{
		select {
		case v := <-numChan:
			if k < 100 {
				fmt.Println(k, v)
				k++
			} else {
				return
			}
		}
	}
}

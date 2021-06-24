package main

import "fmt"

func main() {
	numChan := make(chan int)
	stopChan := make(chan int)

	for i := 0; i < 10; i ++ {
		if i == 9 {
			go func() {
				for j := 0; j < 10; j++ {
					numChan <- j
				}
				stopChan <- 1
			}()
		} else {
			go func() {
				for j := 0; j < 10; j++ {
					numChan <- j
				}
			}()
		}
	}


	for  {
		select {
		case v :=  <- numChan:
			fmt.Println(v)
		case <- stopChan:
			return
		}
	}
}

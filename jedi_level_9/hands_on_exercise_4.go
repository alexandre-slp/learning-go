package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	goRoutines := 100
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(goRoutines)

	fmt.Println(counter)
	for i := 0; i < goRoutines; i++ {
		go func() {
			mutex.Lock()
			c := counter
			c++
			counter = c
			fmt.Println(counter)
			wg.Done()
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

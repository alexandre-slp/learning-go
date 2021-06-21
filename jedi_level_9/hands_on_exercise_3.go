package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	goRoutines := 100
	wg := sync.WaitGroup{}
	wg.Add(goRoutines)
	fmt.Println(counter)
	for i := 0; i < goRoutines; i++ {
		go func() {
			c := counter
			runtime.Gosched()
			c++
			counter = c
			wg.Done()
			fmt.Println(counter)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

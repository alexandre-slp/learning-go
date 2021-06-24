package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	counter := int64(0)
	goRoutines := 100
	wg := sync.WaitGroup{}
	wg.Add(goRoutines)

	fmt.Println(counter)
	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

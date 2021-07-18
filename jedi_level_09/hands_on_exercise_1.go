package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("#CPU", runtime.NumCPU())
	fmt.Println("#Gr", runtime.NumGoroutine())
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("start main")

	go func() {
		fmt.Println("clock")
		wg.Done()
	}()

	go func() {
		fmt.Println("bar")
		wg.Done()
	}()

	fmt.Println("before wait")
	fmt.Println("#CPU", runtime.NumCPU())
	fmt.Println("#Gr", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end main")
	fmt.Println("#CPU", runtime.NumCPU())
	fmt.Println("#Gr", runtime.NumGoroutine())
}

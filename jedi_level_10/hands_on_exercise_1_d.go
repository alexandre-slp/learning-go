package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			c <- n
			fmt.Println("published", n)
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for m := range c {
		fmt.Println(m)
	}
	fmt.Println("finish")
}

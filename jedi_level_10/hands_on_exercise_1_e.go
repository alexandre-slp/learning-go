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
	wg2 := sync.WaitGroup{}

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for m := range c {
			fmt.Println("received", m)
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			fmt.Println("published", n)
			c <- n
		}(i)
	}

	wg.Wait()
	close(c)
	wg2.Wait()
	fmt.Println("finished")
}

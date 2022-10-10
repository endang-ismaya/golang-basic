package main

import (
	"fmt"
	"sync"
)

var counter = 0

func onlyOnce() {
	counter++
}

func main() {
	once := sync.Once{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			once.Do(onlyOnce)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}

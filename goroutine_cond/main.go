package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var wg = sync.WaitGroup{}

func waitCond(value int) {
	defer wg.Done()

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go waitCond(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var gCounter int = 0

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go runAsync(&wg)
	}

	wg.Wait()
	fmt.Println("Program Done!")
}

func runAsync(wg *sync.WaitGroup) {
	defer wg.Done()

	gCounter++
	wg.Add(1)

	fmt.Println("Hello", gCounter)
	time.Sleep(1 * time.Second)

}

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func AddToMap(data *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	data.Store(strconv.Itoa(value), value)
}

func main() {
	var data sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go AddToMap(&data, i, &wg)
	}

	wg.Wait()

	fmt.Println("After GroupWait")
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

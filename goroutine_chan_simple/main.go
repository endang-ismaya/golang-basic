package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel
	ch1 := make(chan string)

	// close
	defer close(ch1)

	// channel can send and received data
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Endang Ismaya" // data need to take by other functions
		fmt.Println("Sending data to ch1 is completed")
	}()

	dataFromCh1 := <-ch1
	fmt.Println("Data from ch1:", dataFromCh1)

	time.Sleep(5 * time.Second)
}

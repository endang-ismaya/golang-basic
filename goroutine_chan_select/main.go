package main

import (
	"fmt"
	"time"
)

var chCounter int = 0

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	defer close(ch1)
	defer close(ch2)

	go giveMeResponse(ch1)
	go giveMeResponse(ch2)

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}

	}

}

func giveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- fmt.Sprintf("chCounter %d", chCounter)
	chCounter++
}

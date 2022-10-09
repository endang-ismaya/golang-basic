package main

import "fmt"

func main() {
	channel := make(chan string, 3)

	channel <- "Data 1"
	channel <- "Data 2"
	channel <- "Data 3"
	close(channel)

	for data := range channel {
		fmt.Println("Menerima: ", data)
	}
}

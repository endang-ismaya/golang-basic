package main

import (
	"fmt"
	"time"
)

func main() {
	testChannelAsParameter()
	testChannelSendOnly()

}

func giveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Endang Ismaya"
}

func testChannelSendOnly() {
	channel := make(chan string)
	defer close(channel)

	go onlySendChannel(channel)
	go onlyReceivedChannel(channel)

	time.Sleep(5 * time.Second)
}

func testChannelAsParameter() {
	channel := make(chan string)
	defer close(channel)

	go giveMeResponse(channel)

	dataFromChannel := <-channel
	fmt.Println("Data from channel", dataFromChannel)

	time.Sleep(5 * time.Second)
}

func onlySendChannel(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Send data to channel"
}

func onlyReceivedChannel(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

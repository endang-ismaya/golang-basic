package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"endangismaya.com/gotest/infinite_loop_with_channel/mylogger"
)

func main() {

	// read input from the user
	// 5 times and write it to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	// execute in go routine for ever
	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something:")

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')

		// assign to channel
		ch <- input

		// wait 1s
		time.Sleep(time.Second)
	}

}

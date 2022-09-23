package main

import (
	"fmt"

	"endangismaya.com/gotest/utils"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message:", message)
		return
	}
	fmt.Println("App End")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error") // this will stop the app
	}
	fmt.Println("App run smoothly")
}

func main() {
	utils.PrintTitle("Test Error")
	runApp(false)

}

package main

import (
	"fmt"

	go_say_hello "github.com/endang-ismaya/go-module-test"
)

func main() {
	hello := go_say_hello.SayHello()
	fmt.Println(hello)

	hello2 := go_say_hello.SayHelloToSomeone("Endang")
	fmt.Println(hello2)

}

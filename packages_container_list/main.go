package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Endang")
	data.PushBack("Ismaya")
	data.PushBack("Wijaya")
	data.PushFront("Mr")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	// iterate linked list
	// front to back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// back to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}

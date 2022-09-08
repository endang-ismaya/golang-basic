package main

import "fmt"

// aggregate types (array, struct)

func main() {
	var myStrings [3]string

	myStrings[0] = "Endang"
	myStrings[1] = "Ismaya"
	myStrings[2] = "Wijaya"

	fmt.Println("First Element in array is", myStrings[0])
	fmt.Println("Second Element in array is", myStrings[1])
	fmt.Println("Third Element in array is", myStrings[2])
}

package main

import (
	"fmt"

	"endangismaya.com/gotest/packageone"
)

var one = "One" // package level variable

func main() {
	var one = "this is a block level variable" // block level variable
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	// function Exported
	packageone.Exported()
}

func myFunc() {
	fmt.Println(one) // using the package level variable
}

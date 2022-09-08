package main

import "fmt"

func main() {
	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2

	// declare type and name and assign value
	var secondNumber = 5

	// declare name, assign value and let GO figure its type
	subtraction := 7

	fmt.Println(firstNumber, secondNumber, subtraction)
}

package main

import "fmt"

// basic types (numbers, strings, booleans)
var myInt int

// var myInt16 int16
// var myInt32 int32
// var myInt64 int64

var myUint uint // only positive values
var myFloat float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 100
	myFloat = 10.1
	myFloat64 = 100.2

	fmt.Println(myInt, myUint, myFloat, myFloat64)

	myString := "" // immutable
	fmt.Println(myString)
	myString = "John"
	fmt.Println(myString)

	var myBool = true
	fmt.Println(myBool)
}

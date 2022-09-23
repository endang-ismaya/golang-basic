package main

import "fmt"

func returnString() interface{} {
	return "Yes a string"
}

func main() {
	var result = returnString()
	var resultString = result.(string) // make sure the return type is string
	fmt.Println(resultString)

	var isAString = returnString()
	switch value := isAString.(type) {
	case string:
		fmt.Println("\"", value, "\" is a string")
	case int:
		fmt.Println("\"", value, "\" is a int")
	default:
		fmt.Println("\"", value, "\" is unknown type")
	}
}

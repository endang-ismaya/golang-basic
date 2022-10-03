package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean) // true
	} else {
		fmt.Println(err.Error())
	}

	// parse int
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number) // 1000000
	} else {
		fmt.Println(err.Error())
	}

	// parse to string
	value := strconv.FormatInt(1000_000, 10)
	fmt.Println(value)        // 1000000
	fmt.Printf("%T\n", value) // string

	// parst alphabet to int
	valueInt, _ := strconv.Atoi("2000")
	fmt.Println(valueInt) // 2000

}

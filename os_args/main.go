package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// go run main.go linux windows macs
	// fmt.Println("os.Args", os.Args)
	// fmt.Println("Path", os.Args[0])
	// fmt.Println("1st argument", os.Args[1])
	// fmt.Println("2nd argument", os.Args[2])
	// fmt.Println("No of items inside os.Args", len(os.Args))

	// > go run main.go 500
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(result)            // 500
	fmt.Printf("%T\n", os.Args[1]) // string
	fmt.Printf("%T\n", result)     // float64
	_ = err                        // mute error
}

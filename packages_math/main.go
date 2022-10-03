package main

import (
	"fmt"
	"math"
)

func main() {
	// round
	fmt.Println(math.Round(10.25)) // 10
	fmt.Println(math.Round(10.73)) // 11

	// floor
	fmt.Println(math.Floor(10.73)) // 10

	// ceil
	fmt.Println(math.Ceil(10.23)) // 11

	// max, min
	fmt.Println(math.Max(10, 20)) // 20
	fmt.Println(math.Min(10, 20)) // 10
}

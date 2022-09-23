package slices1

import (
	"fmt"

	"endangismaya.com/gotest/utils"
)

func CopySlicesWithoutShared() {

	utils.PrintTitle("Copying Slice")

	// appending slices
	// to make a copy of a slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := append([]string{}, cars[0:2]...)
	fmt.Printf("cars (before): %v\n", cars)       // cars (before): [Ford Honda Audi Range Rover]
	fmt.Printf("newCars (before): %v\n", newCars) // newCars (before): [Ford Honda]

	// change newCars
	newCars[0] = "Toyota"
	fmt.Printf("cars (after): %v\n", cars)       // cars (after): [Ford Honda Audi Range Rover]
	fmt.Printf("newCars (after): %v\n", newCars) // newCars (after): [Toyota Honda]

	// Another way to copy a slice
	// copy slice
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // [10 20 30] [10 20 30] 3
}

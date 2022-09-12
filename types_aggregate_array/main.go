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

	// another declaring array
	var a1 = [4]float64{10.0, 11.1, 12.2, 13.3}
	fmt.Printf("%#v\n", a1)

	// ellipsis operator
	a5 := [...]int{1, 2, 5, 1, -10, 17}
	fmt.Printf("the lenght of a5 is %d\n", len(a5)) // he lenght of a5 is 6

	// array operations
	a1[0] = 10.5
	fmt.Printf("%#v\n", a1) // [4]float64{10.5, 11.1, 12.2, 13.3}

	// iterate over an array
	for idx, item := range myStrings {
		fmt.Println(idx, item)
	}
}

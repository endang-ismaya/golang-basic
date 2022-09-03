package main

import "fmt"

func main() {
	// Printf
	// %d -> decimal
	fmt.Printf("My age is %d \n", 34) // out: My age is 34

	// %f -> float
	fmt.Printf("x is %d, y is %.2f \n", 5, 6.6) // out: x is 5, y is 6.60

	// %s -> string
	fmt.Printf("The diameter of a %s with a radius of %d is %f \n", "Circle", 6, float64(6)*2*3.14159) // out: The diameter of a Circle with a radius of 6 is 37.699080

	// %q -> for quoted string
	fmt.Printf("this has a %q\n", "quote") // out: this has a "quote"

	// %v -> replace by any type
	fmt.Printf("The diameter of a %v with a radius of %v is %v \n", "Circle", 6, float64(6)*2*3.14159) // out: The diameter of a Circle with a radius of 6 is 37.699079999999995

	// %T -> type
	fmt.Printf("the type is %T\n", 5)      // out: the type is int
	fmt.Printf("the type is %T\n", 5.5)    // out: the type is float64
	fmt.Printf("the type is %T\n", "five") // out: the type is string

	// %t -> bool
	fmt.Printf("File closed: %t\n", true) // out: File closed: true

	// %b -> base 2
	fmt.Printf("%b \n", 55) // out: 110111
}

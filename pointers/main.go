package main

import "fmt"

func main() {
	name := "Endang"

	fmt.Println(&name) // 0xc000014250

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with a value of %v\n", ptr, ptr) // ptr is of type *int with a value of 0xc0000180d0
	fmt.Printf("address of x is %p\n", &x)                         // 0xc0000180d0

	// creating a pointer using new() built-in function.
	p := new(int) // it creates a pointer called p that is a pointer to an int type

	x = 100
	p = &x // initializing p

	fmt.Printf("p is of type %T with value %v\n", p, p) // => p is of type *int with value 0xc000014140
	fmt.Printf("address of x is %p\n", &x)              // => address of x is 0xc000016120

	//** THE DEREFERENCING OPERATOR **//

	// * in front of a pointer is called the dereferencing operator
	*p = 90 //equivalent to x = 90 because *p is x
	// x and *p is the same thing.

	fmt.Println(*p)                  // => 90
	fmt.Println("*p == x:", *p == x) // => *p == x: true

	fmt.Println("Value of x:", *p) // => Value of x: 90 , equivalent to fmt.Println(x)

	*p = 10        // If I write *p = 10, this is equivalent to x = 10
	*p = *p / 2    //dividing x through the pointer
	fmt.Println(x) // -> 5

	// In a nutshell:
	// &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
	// *pointer => value  -> and if you have pointer you turn it into value value by using the star operator
}

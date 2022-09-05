package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	const (
		c11 = iota
		c12
		c13
	)

	const (
		north = iota // default = 0
		east
		south
		west
	)

	const (
		a = iota * 2 // 0
		b            // 2
		c            // 4
	)

	// skip
	const (
		x = -(iota + 2) // -2
		_               // -3
		y               // -4
		z               // -4
	)

	fmt.Println(c1, c2, c3, c11, c12, c13)
	fmt.Println(north, east, south, west)
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
}

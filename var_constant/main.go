package main

import "fmt"

func main() {
	// Declare a constant
	const days int = 7
	const pi float64 = 3.14
	const secondsInHour int = 3600

	// multiple
	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(days, pi, secondsInHour, n, m, n1, m1, min1, min2, min2)

	// constant rules
	// 1. Can't change a constant

	// 2. Can't initiate a constant at runtime
	// const power = math.Pow(2, 3)

	// 3. Can't use a variable to initialize a constant
	// t := 5
	// const tc = t

	// 4. Can initiate a const from a len -> len built in compile error
	const l1 = len("Hello")
	fmt.Println(l1)

	// untyped constants
	const a float64 = 5.1 // type constant
	const b = 6.7         // untyped constant

	const c float64 = a * b
	fmt.Println(c)

}

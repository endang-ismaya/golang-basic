package main

import "fmt"

func main() {
	// array
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers) // out: [4]int

	// slice
	var cities = []string{"Tangerang", "London", "Tokyo"}
	fmt.Printf("%T\n", cities) // out: []string

	// map
	balances := map[string]float64{
		"USD": 234.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances) // out: map[string]float64

	// struct
	type Person struct {
		name string
		age  int
	}

	var person Person
	fmt.Printf("%T\n", person) // out: main.Person

	// pointer
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr) // out: ptr is of type *int with the value of 0xc0000180c8

	// function type
	fmt.Printf("%T\n", f) // out: func()

	// var golang string = 'Go'
	// fmt.Println(golang) // more than one character in rune literal
}

func f() {}

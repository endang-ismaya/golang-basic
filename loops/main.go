package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for-ever loop
	// for {
	// do nothing here
	// }
	fmt.Println("for i looping")
	forILoop()

	fmt.Println("while looping")
	whileLoop()

	fmt.Println("while looping 2")
	whileLoop2()
}

func forILoop() {
	// for i
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

}

func whileLoop() {
	rand.Seed(time.Now().UnixNano())
	i := 1_000

	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1

		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "and broke out of loop")
}

func whileLoop2() {
	j := 100
	for j >= 0 {
		fmt.Println(j)
		j--
	}
}

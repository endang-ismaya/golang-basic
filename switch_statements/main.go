package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python")

	case "Go", "golang":
		fmt.Println("Good, Go for it!")

	default:
		fmt.Println("Any other language is ok!")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0:
		fmt.Println("Odd number")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	// missing expression
	// means switch true
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening")

	}
}

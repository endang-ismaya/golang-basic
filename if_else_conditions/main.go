package main

import "fmt"

func main() {

	// In Go there is not such a thing like the Truthiness of a variable.
	// Error:
	// if price {
	//  fmt.Println("We have price!")
	// }

	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	if price < 100 {
		fmt.Println("It is OK")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's too expensive")
	}

}

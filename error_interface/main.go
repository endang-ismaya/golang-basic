package main

import (
	"errors"
	"fmt"
	"log"
)

func division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("can't divide with 0")
	}

	result := num1 / num2
	return result, nil
}

func main() {
	result, err := division(100, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result =", result)
}

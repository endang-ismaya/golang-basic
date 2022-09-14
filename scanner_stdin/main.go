package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	// text := scanner.Text()
	// bytes := scanner.Bytes()

	// fmt.Println("Input text:", text)
	// fmt.Println("Input bytes:", bytes)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "q" {
			fmt.Println("Exiting...")
			break
		}

		fmt.Println("You entered:", text)
	}

	fmt.Println("Message after the loop")
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

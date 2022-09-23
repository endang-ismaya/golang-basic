package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	}
}

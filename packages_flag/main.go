package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "Enter host name")
	var user *string = flag.String("user", "root", "Enter your user name")
	var password *string = flag.String("password", "root", "Enter your password")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)

	// go run main.go -host=10.10.0.0 -user=admin -password=hello

}

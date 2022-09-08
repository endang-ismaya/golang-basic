package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	userName  string
	age       int
	favNumber float64
	ownsPet   bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.userName = readString("What is your name?")
	user.age = readInt("How old are you?")
	user.favNumber = readFloat("WHat is your favorite number?")
	user.ownsPet = readBool("Do you own a per (y/n)?")

	fmt.Printf("\nYour name is %s. You are %d years old. Your favorite number is %.2f. Owns a pet: %t\n", user.userName, user.age, user.favNumber, user.ownsPet)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(ask string) string {
	for {
		fmt.Println(ask)
		prompt()

		userInput := getUserInput()

		if userInput == "" {
			fmt.Println("Please enter a name!")
		} else {
			return userInput
		}

	}
}

func readInt(ask string) int {

	for {
		fmt.Println(ask)
		prompt()

		userInput := getUserInput()

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(ask string) float64 {

	for {
		fmt.Println(ask)
		prompt()

		userInput := getUserInput()

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBool(ask string) bool {

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(ask)

		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			return true
		}

	}
}

func getUserInput() string {
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

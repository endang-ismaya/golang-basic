package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats scissors. (paper + 1) % 3 = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// created three slices of stings, each with exactly three entries
var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You shoud buy a lottery ticket",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}

var drawMessages = []string{
	"Great minds think alike",
	"Uh on. Try again!",
	"Nobody wins, but you can try it again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"

	case PAPER:
		computerChoice = "Computer chose PAPER"

	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"

	default:
	}

	// generate a random number from 0-2, which we use to pick a random message
	messageInt := rand.Intn(3)
	// declare a var to hold the message
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		message = loseMessages[messageInt]
	}

	// return winner, computerChoice, roundResult

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}

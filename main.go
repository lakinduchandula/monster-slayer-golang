package main

import (
	"github.com/lakinduchandula/monster-slayer-golang/interaction"
)

var currentRound = 0 // store the current-round

func main() {
	// start game
	startGame()

	winner := "" // contains "Player", "Monster", ""

	// execute until winner == "Player" or "Monster"
	for winner == "" {
		// new round and return whether there is a winner("Player", "Monster") or not ("")
		winner = executeRound()

	}

	// end Game
	endGame()
}

// initial greeting message
func startGame() {
	// call interaction pkg and print greeting message
	interaction.PrintGreeting()
}

// this will hold all the logic for user letting to choose "action"
func executeRound() string {
	/* In every 3 round, currentRound is consider as a special round */

	currentRound += 1                     // increment the current-round by 1
	isSpecialRound := currentRound%3 == 0 // check if current round is a special one

	interaction.GetActions(isSpecialRound) // prompt appropriate actions according to round-count
	return "... logic need to define ..."
}

// which will declare the winner and write the log file
func endGame() {}

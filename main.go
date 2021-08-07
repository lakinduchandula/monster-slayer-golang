package main

import (
	"fmt"
	"github.com/lakinduchandula/monster-slayer-golang/actions"
	"github.com/lakinduchandula/monster-slayer-golang/interaction"
)

var currentRound = 0                   // store the current-round
var gameRound []interaction.RoundStats // empty dynamic slice to store roundStats for every round

func main() {
	// start game
	startGame()

	winner := "" // contains "Player", "Monster", "DRAW", ""

	// execute until winner == "Player" or "Monster"
	for winner == "" {
		// new round and return whether there is a winner("Player", "Monster") or not ("")
		winner = executeRound()
	}

	// end Game
	endGame(winner)
}

// initial greeting message
func startGame() {
	// call interaction pkg and print greeting message
	interaction.PrintGreeting()
}

// this will hold all the logic for user letting to choose "actions"
func executeRound() string {
	// variable declaration for round Statistics
	var playerAttackDmgValue int // default value for int is 0
	var monsterAttackValue int
	var playerHealValue int

	var roundStats *interaction.RoundStats

	/* In every 3 round, currentRound is consider as a special round */
	currentRound += 1                     // increment the current-round by 1
	isSpecialRound := currentRound%3 == 0 // check if current round is a special one

	interaction.GetActions(isSpecialRound, currentRound) // prompt appropriate actions according to round-count
	playerChoice := interaction.GetPlayerChoice(isSpecialRound)

	switch playerChoice {
	case "ATTACK":
		playerAttackDmgValue = actions.AttackMonster(false)
	case "HEAL":
		playerHealValue = actions.HealPlayer()
	case "SPECIAL_ATTACK":
		playerAttackDmgValue = actions.AttackMonster(true)
	default:
		fmt.Println("Error Occurred!")
	}

	// for every round after player attack to monster, monster should attack back
	monsterAttackValue = actions.AttackPlayer()

	// get player-monster current health to judge winner
	playerHealth, monsterHealth := actions.PlayerMonsterHealth()

	// pass the values to NewRoundStats func in output
	roundStats = interaction.NewRoundStats(playerChoice, playerAttackDmgValue, monsterAttackValue, playerHealValue,
		playerHealth, monsterHealth)

	// append new roundStats for gameRound slice
	gameRound = append(gameRound, *roundStats)

	// this will print the current round stats to console
	roundStats.PromptRoundStats()

	// logic to decide winner
	if playerHealth <= 0 && monsterHealth > 0 {
		return "Monster"
	} else if playerHealth > 0 && monsterHealth <= 0 {
		return "Player"
	} else if playerHealth == 0 && monsterHealth == 0 {
		return "DRAW"
	} else if playerHealth < 0 && monsterHealth < 0 {
		if playerHealth > monsterHealth {
			return "Player"
		} else {
			return "Monster"
		}
	}
	return ""
}

// which will declare the winner and write the log file
func endGame(winner string) {
	interaction.DecideWinner(winner)
	interaction.WriteLogFile(&gameRound)
}

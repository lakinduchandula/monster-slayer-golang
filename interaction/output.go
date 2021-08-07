package interaction

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
)

// RoundStats struct for round statics
type RoundStats struct {
	Action               string
	PlayerAttackDmgValue int
	MonsterAttackValue   int
	PlayerHealValue      int
	PlayerCurrentHealth  int
	MonsterCurrentHealth int
}

func NewRoundStats(action string, playerAttackDmgValue int, monsterAttackValue int, playerHealValue int,
	monsterCurrentHealth int, playerCurrentHealth int) *RoundStats {
	newRoundStats := RoundStats{
		action, playerAttackDmgValue, monsterAttackValue,
		playerHealValue, monsterCurrentHealth, playerCurrentHealth,
	}
	return &newRoundStats
}

func PrintGreeting() {
	asciiFigure := figure.NewColorFigure("MONSTER-SLAYER", "", "green", true)
	fmt.Println(asciiFigure)
	fmt.Println("Let's start a new game! =========")
}

func GetActions(isSpecialRound bool, currentRound int) {
	// only executes for new rounds, not for first round
	if currentRound != 1 {
		fmt.Printf("\n----------- NEW ROUND -----------\n")
	}

	fmt.Println("===== Pick an Action Below! =====")
	fmt.Println("\t- (1) Attack Monster")
	fmt.Println("\t- (2) Heal My Self")

	// only display in a special round
	if isSpecialRound {
		fmt.Println("\t- (3) Special Attack")
	}
}

func DecideWinner(winner string) {
	fmt.Printf("\n--------- Game is Over ----------\n")

	if winner == "DRAW" {
		fmt.Println("\t The game is Draw")
	} else {
		fmt.Printf("\t%s has won the game.\n", winner)
	}

	fmt.Println("Thanks for playing!!  ===========")
}

// PromptRoundStats is a method of roundStats struct
func (stats *RoundStats) PromptRoundStats() {
	{
		fmt.Println("\n ###### Latest Statistics ###### ")

		if stats.Action == "SPECIAL_ATTACK" {
			fmt.Println(" SELECTED CHOICE >>> SP.(ATTACK)")
		} else {
			fmt.Println("  SELECTED CHOICE  >>> ", stats.Action)
		}
		fmt.Printf("%16s%13s\n", "Player", "Monster")

		fmt.Printf("%s%11d%13d\n", "DMG", stats.PlayerAttackDmgValue, stats.MonsterAttackValue)
		if stats.Action == "HEAL" {
			fmt.Printf("%s%10d%13s\n", "HEAL", stats.PlayerHealValue, "-")
		} else {
			fmt.Printf("%s%10d%13s\n", "HEAL", stats.PlayerHealValue, "-")
		}
		fmt.Printf("%s%6d%13d\n", "C-HEALTH", stats.PlayerCurrentHealth, stats.MonsterCurrentHealth)
	}
}

func WriteLogFile(rounds *[]RoundStats) {
	file, err := os.Create("gameLog.txt")

	if err != nil {
		fmt.Println("Saving to a log file is failed")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                  fmt.Sprint(index + 1),
			"Action":                 value.Action,
			"Player Attack Damage":   fmt.Sprint(value.PlayerAttackDmgValue),
			"Player Heal Value":      fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage":  fmt.Sprint(value.MonsterAttackValue),
			"Player Current Health":  fmt.Sprint(value.PlayerCurrentHealth),
			"Monster Current Health": fmt.Sprint(value.MonsterCurrentHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Error while writing log file")
			continue
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("file was unable to close")
		return
	}
	fmt.Println("Writing log file was successful")
}

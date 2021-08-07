package interaction

import "fmt"

// RoundStats struct for round statics
type RoundStats struct {
	Action               string
	PlayerAttackDmgValue int
	MonsterAttackValue   int
	PlayerHealValue      int
	PlayerCurrentHealth  int
	MonsterCurrentHealth int
}

func PrintGreeting() {
	fmt.Println("\t\tMONSTER-SLAYER")
	fmt.Println("Let's start a new game! =========")
}

func GetActions(isSpecialRound bool, currentRound int) {
	// only executes for new rounds, not for first round
	if currentRound != 1 {
		fmt.Printf("----------- NEW ROUND -----------\n\n")
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

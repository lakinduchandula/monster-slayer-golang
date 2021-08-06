package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER-SLAYER")
	fmt.Println("Let's start a new game! ========")
}

func GetActions(isSpecialRound bool) {
	fmt.Println("===== Pick an Action Below! =====")
	fmt.Println("")
	fmt.Println("\t- (1) Attack Monster")
	fmt.Println("\t- (2) Heal My Self")

	// only display in a special round
	if isSpecialRound {
		fmt.Println("\t- (3) Special Attack")
	}
}

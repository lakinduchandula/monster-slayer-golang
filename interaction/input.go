package interaction

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialAttack bool) string {
	// infinite loop
	for {
		userInput, _ := getPlayerInput()

		switch userInput {
		case "1":
			return "ATTACK"
		case "2":
			return "HEAL"
		case "3":
			if isSpecialAttack {
				return "SPECIAL_ATTACK"
			} else {
				fmt.Println("Invalid input. Please try again.")
			}
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Enter Your Choice : ")

	// read the user-input
	playerInput, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("input is invalid")
	}

	playerInput = strings.Replace(playerInput, "\n", "", -1)

	return playerInput, nil
}

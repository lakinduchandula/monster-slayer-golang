package main

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
func startGame() {}

// this will hold all the logic for user letting to choose "action"
func executeRound() string {
	return "... logic need to define ..."
}

// which will declare the winner and write the log file
func endGame() {}

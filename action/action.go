package action

import (
	"math/rand"
	"time"
)

// this contains the source value to generate a random number
var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

func AttackMonster() {}

func HealPlayer() {}

func AttackPlayer() {}

func generateRandomBetween(min int, max int) int {
	return randomGenerator.Intn(max-min) + min
}

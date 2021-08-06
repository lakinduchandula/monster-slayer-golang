package action

import (
	"math/rand"
	"time"
)

// this contains the source value to generate a random number
var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

// player-monster health values
var currentPlayerHealth = 100
var currentMonsterHealth = 100

func AttackMonster(isSpecialAttack bool) {
	// min-max attack value for general attack
	minAttackValue := 5
	maxAttackValue := 10

	if isSpecialAttack {
		minAttackValue = 10
		maxAttackValue = 20
	}

	// calculate damage-value
	damageValue := generateRandomBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= damageValue // deduct damageValue from currentMonsterHealth

}

func HealPlayer() {
	// heal values for general case
	minHealValue := 7
	maxHealValue := 15

	// calculate heal-value
	healValue := generateRandomBetween(minHealValue, maxHealValue)
	currentPlayerHealth += healValue // current heal value

	// system won't allow playerHealth to be greater than 100
	if currentPlayerHealth > 100 {
		currentPlayerHealth = 100
	}
}

func AttackPlayer() {
	minAttackValue := 8
	maxAttackValue := 15

	// calculate attack-player value
	playerAttackValue := generateRandomBetween(minAttackValue, maxAttackValue)

	// update current-player-health
	currentPlayerHealth -= playerAttackValue
}

func generateRandomBetween(min int, max int) int {
	return randomGenerator.Intn(max-min) + min
}

package actions

import (
	"math/rand"
	"time"
)

// this contains the source value to generate a random number
var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

// player-monster health values
var currentPlayerHealth = PLAYER_HEALTH
var currentMonsterHealth = MONSTER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	// min-max attack value for general attack
	minAttackValue := PLAYER_ATTACK_MIN_DMG_GENERAL
	maxAttackValue := PLAYER_ATTACK_MAX_DMG_GENERAL

	if isSpecialAttack {
		minAttackValue = PLAYER_ATTACK_MIN_DMG_SPECIAL
		maxAttackValue = PLAYER_ATTACK_MAX_DMG_SPECIAL
	}

	// calculate damage-value
	damageValue := generateRandomBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= damageValue // deduct damageValue from currentMonsterHealth

	return damageValue
}

func HealPlayer() int {
	// heal values for general case
	minHealValue := PLAYER_MIN_HEAL_VALUE
	maxHealValue := PLAYER_MAX_HEAL_VALUE

	// calculate heal-value
	healValue := generateRandomBetween(minHealValue, maxHealValue)
	currentPlayerHealth += healValue // current heal value

	// system won't allow playerHealth to be greater than 100
	if currentPlayerHealth > PLAYER_HEALTH {
		currentPlayerHealth = PLAYER_HEALTH
	}

	return healValue
}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	// calculate attack-player value
	playerAttackValue := generateRandomBetween(minAttackValue, maxAttackValue)

	// update current-player-health
	currentPlayerHealth -= playerAttackValue

	return playerAttackValue
}

func PlayerMonsterHealth() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandomBetween(min int, max int) int {
	return randomGenerator.Intn(max-min) + min
}

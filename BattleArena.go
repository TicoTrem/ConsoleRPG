package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var baseXPToGrant int = 10

// to be in a battle, they must be able to defend
func StartBattle(MC MainCharacter, enemies []Enemy) bool {

	enemiesCopy := make([]Enemy, len(enemies))
	copy(enemiesCopy, enemies)

	// delete by index
	enemies = append(enemies[:0], enemies[0+1:]...)

	for len(enemies) > 0 && !MC.IsDead() {
		time.Sleep(1 * time.Second)
		mcTakeTurn(MC, enemies)
	}

	for i := 0; i < len(enemies); i++ {
		if !enemies[i].IsDead() && !MC.IsDead() {
			time.Sleep(1 * time.Second)
			enemies[i].Attack(&MC)
		}
	}

	MCWon := false
	if !MC.IsDead() {
		grantExperience(MC, enemiesCopy)
		MCWon = true
	}

	return MCWon
}

func mcTakeTurn(MC MainCharacter, enemies []Enemy) {
	displayEnemies(enemies)
	var chosenEnemy *Enemy
	for {
		var enemySelection string
		fmt.Println("Select an enemy to attack (Enter their number)")
		fmt.Scanln(&enemySelection)
		num, err := strconv.Atoi(enemySelection)
		if err != nil || num > len(enemies) || num < 1 {
			fmt.Println("Invalid Selection!")
			continue
		}
		// dereference before indexing
		chosenEnemy = &enemies[num-1]
		MC.Attack(chosenEnemy)
		fmt.Println(chosenEnemy.HP)
		break
	}

	if chosenEnemy.IsDead() {
		fmt.Println("It was a killing blow!")
		fmt.Println(rand.Intn(len(MC.CatchPhrases)))
	}

}

func displayEnemies(enemies []Enemy) {
	for index, enemy := range enemies {
		fmt.Printf("Enemy %v: %v HP %v/%v Weapon %v\n", index+1, enemy.CharacterName, enemy.HP, enemy.MaxHP, enemy.weapon.Name)
	}
}

func grantExperience(MC MainCharacter, originalEnemies []Enemy) {

	for _, enemy := range originalEnemies {
		MC.GrantXP(enemy.Level * baseXPToGrant)
	}
}

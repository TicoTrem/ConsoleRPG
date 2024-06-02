package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var baseXPToGrant int = 10

// to be in a battle, they must be able to defend
func StartBattle(MC *MainCharacter, enemies []Enemy) bool {

	enemiesCopy := make([]Enemy, len(enemies))
	copy(enemiesCopy, enemies)

	for len(enemies) > 0 && !MC.IsDead() {
		time.Sleep(1 * time.Second)
		mcTakeTurn(MC, &enemies)

		for i := 0; i < len(enemies); i++ {
			enemyTakeTurn(&enemies[i], MC)
		}
	}

	MCWon := false
	if !MC.IsDead() {
		grantExperience(MC, enemiesCopy)
		MCWon = true
	}

	return MCWon
}

func enemyTakeTurn(enemy *Enemy, MC *MainCharacter) {
	if !MC.IsDead() {
		if !enemy.IsDead() {
			time.Sleep(1 * time.Second)
			enemy.Attack(MC)
		}
	} else {
		fmt.Printf("%v has fallen to %v!\n", MC.CharacterName, enemy.CharacterName)
		enemy.SayRandomCatchphrase()
	}
}

func mcTakeTurn(MC *MainCharacter, enemies *[]Enemy) {
	displayEnemies(*enemies)
	var chosenEnemy *Enemy
	for {
		var enemySelection string
		fmt.Println("Select an enemy to attack (Enter their number)")
		fmt.Scanln(&enemySelection)
		num, err := strconv.Atoi(enemySelection)
		if err != nil || num > len(*enemies) || num < 1 {
			fmt.Println("Invalid Selection!")
			continue
		}
		// dereference before indexing
		chosenEnemy = &((*enemies)[num-1])
		MC.Attack(chosenEnemy)
		break
	}

	if chosenEnemy.IsDead() {
		killEnemy(enemies, chosenEnemy)
		MC.SayRandomCatchphrase()
	}
}

func killEnemy(enemies *[]Enemy, enemyToRemove *Enemy) bool {

	for i, v := range *enemies {
		if reflect.DeepEqual(v, *enemyToRemove) {
			*enemies = append((*enemies)[:i], (*enemies)[i+1:]...)
			fmt.Println("It was a killing blow!")
			return true
		}
	}
	panic("Tried to kill an enemy that is not in the enemies list")
}

func displayEnemies(enemies []Enemy) {
	for index, enemy := range enemies {
		fmt.Printf("Enemy %v: %v HP %v/%v Weapon %v\n", index+1, enemy.CharacterName, enemy.HP, enemy.MaxHP, enemy.weapon.Name)
	}
}

func grantExperience(MC *MainCharacter, originalEnemies []Enemy) {

	for _, enemy := range originalEnemies {
		MC.GrantXP(enemy.Level * baseXPToGrant)
	}
}

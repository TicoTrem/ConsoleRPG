package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(5))
	fmt.Println("Hello adventurer. What is your name?")
	var userName string
	fmt.Scanln(&userName)

	MC := NewMainCharacter(userName, *NewTrainingSword())
	fmt.Printf("Well then %v, lets get you on your way.\n", MC.CharacterName)
	fmt.Println("We are in dire need of common swordsmen, and you are not special. Take this.")

	DisplaySystemMessage(fmt.Sprintf("You obtained a %v!", MC.weapon.Name))

	StartBattle(*MC, []Enemy{
		{CharacterComponent: CharacterComponent{CharacterName: "Neil", CatchPhrases: []string{"Neil Doyleee!!!!"}}, EnemyAttackComponent: EnemyAttackComponent{weapon: *NewFists()}, EnemyDefendComponent: EnemyDefendComponent{MaxHP: 20, HP: 20, Armour: 0, DodgeChance: 0.05}},
		{CharacterComponent: CharacterComponent{CharacterName: "Seneca", CatchPhrases: []string{"Neil Doyleee!!!!"}}, EnemyAttackComponent: EnemyAttackComponent{weapon: *NewFists()}, EnemyDefendComponent: EnemyDefendComponent{MaxHP: 20, HP: 20, Armour: 0, DodgeChance: 0.05}},
		{CharacterComponent: CharacterComponent{CharacterName: "Doyle", CatchPhrases: []string{"Neil Doyleee!!!!"}}, EnemyAttackComponent: EnemyAttackComponent{weapon: *NewFists()}, EnemyDefendComponent: EnemyDefendComponent{MaxHP: 20, HP: 20, Armour: 0, DodgeChance: 0.05}},
	})
}

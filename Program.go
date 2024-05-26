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
		{Character: Character{CharacterName: "John", CatchPhrases: []string{"Neil Doyle!!!!!"}}, CombatCharacter: CombatCharacter{HP: 20, MaxHP: 20, weapon: *NewFists(), DodgeChance: 0.05, Level: 1}},
		{Character: Character{CharacterName: "Tony", CatchPhrases: []string{"Neil Doyle!!!!!"}}, CombatCharacter: CombatCharacter{HP: 20, MaxHP: 20, weapon: *NewFists(), DodgeChance: 0.05, Level: 1}},
		{Character: Character{CharacterName: "Jessica", CatchPhrases: []string{"Neil Doyle!!!!!"}}, CombatCharacter: CombatCharacter{HP: 20, MaxHP: 20, weapon: *NewFists(), DodgeChance: 0.05, Level: 1}}})
}

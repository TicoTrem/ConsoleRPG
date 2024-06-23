package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello adventurer. What is your name?")
	var userName string
	fmt.Scanln(&userName)

	MC := NewMainCharacter(userName, *NewTrainingSword())
	fmt.Printf("Well then %v, lets get you on your way.\n", MC.CharacterName)
	fmt.Println("We are in dire need of common swordsmen, and you are not special. Take this.")

	DisplaySystemMessage(fmt.Sprintf("You obtained a %v!", MC.weapon.Name))

	StartBattle(MC, []Enemy{
		*NewBozo(*NewFists(), 1),
	})
	MC.DisplayCharacterInfo()

	StartBattle(MC, []Enemy{
		*NewBozo(*NewFists(), 1),
	})

	MC.DisplayCharacterInfo()
	time.Sleep(60 * time.Second)
}

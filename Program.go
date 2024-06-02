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
	fmt.Println("HI ANNA! This is the current end of the game, it is not much right now, but the systems" +
		"in place to make it highly expandable are extensive!")
	time.Sleep(60 * time.Second)
}

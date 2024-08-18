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
		// *NewRainbowUnicorn(*NewFists(), 2),
	})
	MC.DisplayCharacterInfo()

	fmt.Println("Fantastic work, that would've been embarassing if you lost...")
	fmt.Println("Well we don't have enough time to explain EVERYTHING to EVERYONE.")
	fmt.Println("Go down one of these paths and you should end up at a small settlement.")
	fmt.Println("Look for a woman named Giji. Bye.")

	time.Sleep(60 * time.Second)
}

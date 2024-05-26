package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello adventurer. What is your name?")
	var userName string
	fmt.Scanln(&userName)
	MC := NewMainCharacter(userName, Weapon{})
	fmt.Printf("Well then %v, lets get you on your way.\n", MC.CharacterName)
	fmt.Println("We are in dire need of common swordsman, and you are not special. Take this.")

	DisplaySystemMessage("You obtained a training sword!")
}

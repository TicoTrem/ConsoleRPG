package main

import (
	"fmt"
)

type CharacterComponent struct {
	CharacterName string
	CatchPhrases  []string
}

func (c CharacterComponent) Say(msg string) {
	fmt.Printf("%v: %v", c.CharacterName, msg)
}

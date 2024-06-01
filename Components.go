package main

import (
	"fmt"
	"math/rand"
)

type CharacterComponent struct {
	CharacterName string
	CatchPhrases  []string
}

func (c CharacterComponent) Say(msg string) {
	fmt.Printf("%v: %v\n", c.CharacterName, msg)
}

func (c CharacterComponent) SayRandomCatchphrase() {
	c.Say(c.CatchPhrases[rand.Intn(len(c.CatchPhrases))])
}

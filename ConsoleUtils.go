package main

import (
	"fmt"
)

func DisplaySystemMessage(msg string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("* %v *\n", msg)
	fmt.Println("--------------------------------------------------------------------------------------")
}

func ChoosePath(options []string, functions []func()) {
	fmt.Println("There are multiple paths you could take, choose wisely...")
	PromptAndChoose(options, functions)
}

func PromptAndChoose(options []string, functions []func()) {
	for i := 0; i < len(options); i++ {

	}
}

// chest rooms and battles implement Event interface, so they can all be here
// Event has a GetType method
func ConstructPathPrompt(events []Event) {
	var prompt string = "Would you like to "
	for index, value := range events {
		var bLast bool = false
		if index == len(events)-1 {
			bLast = true
			prompt += "or"
		}
		switch value.getEventType() {
		case "ChestRoom":
			prompt += "venture into the gamble that is the chest room"
		case "Battle":
			prompt += "try your hand at fighting some baddies"
		}
		if bLast {
			prompt += "."
		} else {
			prompt += ","
		}

	}
}

type Event struct {
	eventType string
}

func (e *Event) getEventType() string {
	return e.eventType
}

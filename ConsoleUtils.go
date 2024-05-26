package main

import (
	"fmt"
)

func DisplaySystemMessage(msg string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("* %v *\n", msg)
	fmt.Println("--------------------------------------------------------------------------------------")
}

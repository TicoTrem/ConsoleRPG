package main

import (
	"fmt"
	"math/rand"
)

// use embedding for code sharing, not for composition like with DefendComponent and stuff
// we have Enemy and MainCharacter both with Level but we do not repeat code so we do not
// care
type MainCharacter struct {
	CharacterComponent
	MCAttackComponent
	EnemyDefendComponent
	MCLevelComponent
	XP         int
	CritChance float32
}

func (c MainCharacter) GetName() string {
	return c.CharacterName
}
func (c MainCharacter) GetStatMultiplier() float32 {
	return 1 + ((float32(c.Level) * 0.1) - 0.1)
}
func (c MainCharacter) GetCritChance() float32 {
	return c.CritChance + c.weapon.CritChance
}

func NewMainCharacter(name string, weapon Weapon) *MainCharacter {
	if name == "" {
		panic("Name is required!")
	}

	if weapon.Name == "" {
		// weapon not set, set to default to fists
		weapon = *NewFists()
	}

	character := &MainCharacter{
		CharacterComponent:   CharacterComponent{CharacterName: name, CatchPhrases: []string{"Yahoooo", "get rekt kid", "pfft Bozo"}},
		MCAttackComponent:    MCAttackComponent{weapon: weapon},
		EnemyDefendComponent: EnemyDefendComponent{MaxHP: 20, HP: 20, Armour: 0, DodgeChance: 0.05},
		MCLevelComponent:     MCLevelComponent{Level: 1, XP: 0},
	}

	// now pass the character as the object that satisfies the CharacterData interface
	character.EnemyDefendComponent.Named = character
	character.EnemyDefendComponent.StatMultiplier = character
	character.MCAttackComponent.Named = character
	character.MCAttackComponent.Criticaller = character
	character.MCAttackComponent.StatMultiplier = character

	return character

}

type MCAttackComponent struct {
	weapon Weapon
	Named
	StatMultiplier
	Criticaller // we are basically saying to have this component or implementation of the interface, you
	// also need all of the other components that will allow you to satisfy this interface
}

// now we can attack anything that can take damage
func (m MCAttackComponent) Attack(d IDefend) {
	isCritical := m.isCritical(m.weapon)

	damageToDeal := m.calculateDamage(m.weapon, isCritical)
	bHit, str := d.TakeDamage(damageToDeal, true)
	if bHit {
		fmt.Printf("%v used their %v and dealt %v damage.\n"+str,
			m.GetName(),
			m.weapon.Name,
			damageToDeal)

		if isCritical {
			fmt.Println("Critical Hit!")
		}
	}

}

func (m MCAttackComponent) calculateDamage(weapon Weapon, isCritical bool) int {
	damage := int(float32(weapon.BaseDamage) * m.GetStatMultiplier())

	if isCritical {
		damage *= int(m.weapon.CriticalBonus)
	}

	return damage
}

func (m MCAttackComponent) isCritical(w Weapon) bool {

	rand := rand.Float32()
	// if the random float is less than crit chance, it is a critical hit
	return rand < m.GetCritChance()
}

type MCLevelComponent struct {
	Level int
	XP    int
}

func (l MCLevelComponent) levelUp() {
	l.Level++
	l.XP = 0
}

func (l MCLevelComponent) GrantXP(nXP int) {
	l.XP += nXP
	if l.XP > 100 {
		l.levelUp()
		l.XP = 0
	}
}

package main

import (
	"fmt"
	"math/rand"
)

type Character struct {
	CharacterName string
	CatchPhrases  []string
}

func (c Character) Say(msg string) {
	fmt.Printf("%v: %v", c.CharacterName, msg)
}

// not saying "pick a DefendModule implementation", instead it is saying
// "This has all the methods and attributes of this DefendModule, which is an implementation without an interface"
// If it had an interface, we could specify a different implementation of the Defend() method for each different ComatCharacter

// If i don't use interfaces, we can only choose different values for the attributes when we make a new weapon
// If I use interfaces, we can add different implementations of things like the attack method, and even add helper methods for it

// simplifying by putting all identical code between MainCharacter
// and Enemy in to the CombatCharacter struct, if there is anything not
// identical, don't think about it, just add it to the enemy class with its
// own implementation

// ALSO I'm planning on just not using some featuers like interfaces until I find a use for them
// instead of forcing them where I think they would be the best

// not storing a pointer in the struct because
// the struct should own and not share this copy of the weapon

type CombatCharacter struct {
	Character
	DefendModule
	HP          int
	MaxHP       int
	DodgeChance float32
	Armour      int
	Level       int
	weapon      Weapon
}

func (c CombatCharacter) IsDead() bool {
	return c.HP < 0
}

type MainCharacter struct {
	CombatCharacter
	XP              int
	NextAttackBonus int
	CritChance      float32
}

func NewMainCharacter(name string, weapon Weapon) *MainCharacter {
	if name == "" {
		panic("Name is required!")
	}

	if weapon.Name == "" {
		// weapon not set, set to default to fists
		weapon = *NewFists()
	}

	return &MainCharacter{
		CombatCharacter: CombatCharacter{
			Character:   Character{CharacterName: name},
			MaxHP:       20,
			HP:          20,
			DodgeChance: 0.05,
			Level:       1,
			weapon:      weapon,
		},
	}

}

// now we can attack anything that can take damage
func (m MainCharacter) Attack(d IDefend) {
	isCritical := m.IsCritical(m.weapon)

	damageToDeal := m.CalculateDamage(m.weapon, isCritical)
	d.TakeDamage(damageToDeal, true)
	fmt.Printf("%v used their %v and dealt %v damage to %v\n%v now has %v/%v health",
		m.CharacterName,
		m.weapon.Name,
		damageToDeal,
		m.CharacterName,
		m.CharacterName,
		m.HP, m.MaxHP)

	if isCritical {
		fmt.Println("Critical Hit!")
	}

}

func (m MainCharacter) CalculateDamage(weapon Weapon, isCritical bool) int {
	damage := int(float32(weapon.BaseDamage+m.NextAttackBonus) * m.StatMultiplier())

	if isCritical {
		damage *= int(m.weapon.CriticalBonus)
	}

	return damage
}

func (m MainCharacter) IsCritical(w Weapon) bool {
	critChance := w.CritChance + m.CritChance

	rand := rand.Float32()
	// if the random float is less than crit chance, it is a critical hit
	return rand < critChance
}

func (m MainCharacter) GrantXP(xpToGrant int) {
	m.XP += xpToGrant
	if m.XP > 100 {
		m.LevelUp()
		m.XP = 0
	}
}

func (m MainCharacter) LevelUp() {
	m.Level++
}

func (m MainCharacter) StatMultiplier() float32 {
	return 1 + ((float32(m.Level) * 0.1) - 0.1)
}

type Enemy struct {
	CombatCharacter
}

type IDefend interface {
	TakeDamage(nDamage int, bBlockable bool)
}

// because we have a common implementation we want to share
// not just using the interface, because we might eventually want the character
// and the objects to have a different take damage implementation
// This implementation uses an HP so we have a struct with HP in it as well
type DefendModule struct {
	HP     int
	Armour int
}

func (d DefendModule) TakeDamage(nDamage int, bBlockable bool) {
	if bBlockable {
		d.HP -= nDamage - d.Armour
	} else {
		d.HP -= nDamage
	}
}

// could have just added a TakeDamage method to this AttackableObject struct
// but then we would be repeating code, so make an implementation and add it to all
type AttackableObject struct {
	DefendModule
	Name   string
	HP     int
	Armour int
}

func NewTree(health int) *AttackableObject {
	return &AttackableObject{
		Name:   "Tree",
		HP:     health,
		Armour: 1,
	}
}

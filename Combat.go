package main

import (
	"fmt"
	"math/rand"
)

type DamageType int

const (
	Regular DamageType = iota
	Poison
	Bleeding
	Electric
)

type Weapon struct {
	Name             string
	BaseDamage       int
	WeaponDamageType DamageType
	DescriptorWord   string
	AttackMessage    string
	CritChance       float32
	CriticalBonus    float32
}

type Character struct {
	CharacterName string
	CatchPhrases  []string
}

func (c Character) Say(msg string) {
	fmt.Printf("%v: %v", c.CharacterName, msg)
}

type IDefend interface {
	TakeDamage(damageToTake int, isBlockable bool)
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
type CombatCharacter struct {
	Character
	HP          int
	MaxHP       int
	DodgeChance float32
	Armour      int
	Level       int
	weapon      Weapon
}



func NewMainCharacter(name string, weapon Weapon) *MainCharacter {
	if name == "" {
		panic("Name is required!")
	}

	if weapon.Name == "" {
		// weapon not set, set to default o ffists
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

func (c CombatCharacter) IsDead() bool {
	return c.HP < 0
}

type Enemy struct {
	CombatCharacter
}

// Go doesnt have the thing where an if an embedded thingy satisfies an interface,
// you can assign that object to the interface type, this means that it was never meant
// to have the level of abstraction where both MainCharacter and Enemy implemented a defend
// module with different implementations
func (e Enemy) TakeDamage(nDamage int, bBlockable bool) {
	if bBlockable {
		e.HP -= nDamage - e.Armour
	} else {
		e.HP -= nDamage
	}
}

type MainCharacter struct {
	CombatCharacter
	XP              int
	NextAttackBonus int
	CritChance      float32
}

func (m MainCharacter) Attack(e Enemy) {
	isCritical := m.IsCritical(m.weapon)

	damageToDeal := m.CalculateDamage(m.weapon, isCritical)
	e.TakeDamage(damageToDeal, true)
	fmt.Printf("%v used their %v and dealt %v damage to %v\n%v now has %v/%v health",
		m.CharacterName,
		m.weapon.Name,
		damageToDeal,
		e.CharacterName,
		e.CharacterName,
		e.HP, e.MaxHP)

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

package main

import (
	"fmt"
	"math/rand"
)

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

type IDefend interface {
	TakeDamage(nDamage int, bBlockable bool) (bool, string)
	IsDead() bool
}
type IAttack interface {
	Attack(d IDefend)
}

type CombatCharacter struct {
	Character
	MaxHP       int
	HP          int
	DodgeChance float32
	Armour      int
	Level       int
	weapon      Weapon
}

func (c *CombatCharacter) TakeDamage(nDamage int, bBlockable bool) (bool, string) {
	bHit := true
	var str string
	if rand.Float32() > c.DodgeChance {
		if bBlockable {
			c.HP -= nDamage - c.Armour
		} else {
			c.HP -= nDamage
		}
		str = fmt.Sprintf("%v now has %v/%v health\n",
			c.CharacterName,
			c.HP,
			c.MaxHP)
	} else {
		fmt.Println("Attack dodged, no damage was dealt")
		bHit = false
	}
	return bHit, str
}

func (c CombatCharacter) IsDead() bool {
	if c.HP <= 0 {
		return true
	} else {
		return false
	}
}

func (c CombatCharacter) StatMultiplier() float32 {
	return 1 + ((float32(c.Level) * 0.1) - 0.1)
}

type Character struct {
	CharacterName string
	CatchPhrases  []string
}

func (c Character) Say(msg string) {
	fmt.Printf("%v: %v", c.CharacterName, msg)
}

type Enemy struct {
	CombatCharacter
}

func (e Enemy) Attack(d IDefend) {
	isCritical := e.isCritical(e.weapon)

	damageToDeal := e.calculateDamage(e.weapon, isCritical)
	bHit, str := d.TakeDamage(damageToDeal, true)
	if bHit {
		fmt.Printf("%v used their %v and dealt %v damage.\n"+str,
			e.CharacterName,
			e.weapon.Name,
			damageToDeal)

		if isCritical {
			fmt.Println("Critical Hit!")
		}
	}

}

func (e Enemy) calculateDamage(weapon Weapon, isCritical bool) int {
	damage := int(float32(weapon.BaseDamage) * e.StatMultiplier())

	if isCritical {
		damage *= int(e.weapon.CriticalBonus)
	}

	return damage
}

func (e Enemy) isCritical(w Weapon) bool {
	critChance := w.CritChance

	rand := rand.Float32()
	// if the random float is less than crit chance, it is a critical hit
	return rand < critChance
}

// use embedding for code sharing, not for composition like with DefendComponent and stuff
// we have Enemy and MainCharacter both with Level but we do not repeat code so we do not
// care
type MainCharacter struct {
	CombatCharacter
	XP         int
	CritChance float32
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
		CombatCharacter: CombatCharacter{Character: Character{CharacterName: name, CatchPhrases: []string{"Yahoooo", "get rekt kid", "pfft Bozo"}}, MaxHP: 20, HP: 20, DodgeChance: 0.05, Armour: 0, Level: 1, weapon: weapon},
	}

}

// now we can attack anything that can take damage
func (m MainCharacter) Attack(d IDefend) {
	isCritical := m.isCritical(m.weapon)

	damageToDeal := m.calculateDamage(m.weapon, isCritical)
	bHit, str := d.TakeDamage(damageToDeal, true)
	if bHit {
		fmt.Printf("%v used their %v and dealt %v damage.\n"+str,
			m.CharacterName,
			m.weapon.Name,
			damageToDeal)

		if isCritical {
			fmt.Println("Critical Hit!")
		}
	}

}

func (m MainCharacter) calculateDamage(weapon Weapon, isCritical bool) int {
	damage := int(float32(weapon.BaseDamage) * m.StatMultiplier())

	if isCritical {
		damage *= int(m.weapon.CriticalBonus)
	}

	return damage
}

func (m MainCharacter) isCritical(w Weapon) bool {
	critChance := w.CritChance + m.CritChance

	rand := rand.Float32()
	// if the random float is less than crit chance, it is a critical hit
	return rand < critChance
}

func (m MainCharacter) GrantXP(xpToGrant int) {
	m.XP += xpToGrant
	if m.XP > 100 {
		m.levelUp()
		m.XP = 0
	}
}

func (m MainCharacter) levelUp() {
	m.Level++
}

// could have just added a TakeDamage method to this AttackableObject struct
// but then we would be repeating code, so make an implementation and add it to all
type AttackableObject struct {
	Name   string
	HP     int
	Armour int
}

func (a *AttackableObject) TakeDamage(nDamage int, bBlockable bool) (bool, string) {
	if bBlockable {
		a.HP -= nDamage - a.Armour
	} else {
		a.HP -= nDamage
	}
	str := fmt.Sprintf("%v now has %v health\n",
		a.Name,
		a.HP)
	return true, str
}

func NewTree(health int) *AttackableObject {
	return &AttackableObject{
		Name:   "Tree",
		HP:     health,
		Armour: 1,
	}
}

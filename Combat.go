package main

import (
	"fmt"
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
type ILevel interface {
	levelUp()
	GrantXP(nXP int)
}

type Healthed interface {
	GetMaxHealth() int
	SetMaxHealth(amount int)
}
type Named interface {
	GetName() string
}
type StatMultiplier interface {
	GetStatMultiplier() float32
}

type Criticaller interface {
	GetCritChance() float32
}

// could have just added a TakeDamage method to this AttackableObject struct
// but then we would be repeating code, so make an implementation and add it to all
type AttackableObject struct {
	Name string
	NonAttackingDefendComponent
}

func (a AttackableObject) GetName() string {
	return a.Name
}

type NonAttackingDefendComponent struct {
	HP     int
	Armour int
	Named
}

func (a *NonAttackingDefendComponent) IsDead() bool {
	if a.HP <= 0 {
		return true
	} else {
		return false
	}
}

func (a *NonAttackingDefendComponent) TakeDamage(nDamage int, bBlockable bool) (bool, string) {
	if bBlockable {
		a.HP -= nDamage - a.Armour
	} else {
		a.HP -= nDamage
	}
	str := fmt.Sprintf("%v now has %v health\n",
		a.GetName(),
		a.HP)
	return true, str
}

func NewTree(health int) *AttackableObject {
	tree := &AttackableObject{
		Name:                        "Tree",
		NonAttackingDefendComponent: NonAttackingDefendComponent{HP: 10, Armour: 1},
	}

	tree.Named = tree
	return tree
}

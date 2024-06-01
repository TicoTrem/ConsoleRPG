package main

import (
	"fmt"
	"math/rand"
)

type Enemy struct {
	CharacterComponent
	EnemyAttackComponent
	EnemyDefendComponent
	Level int
}

func (e Enemy) GetStatMultiplier() float32 {
	return 1 + ((float32(e.Level) * 0.1) - 0.1)
}

func (e Enemy) GetCritChance() float32 {
	return e.weapon.CritChance
}

func (e Enemy) GetName() string {
	return e.CharacterName
}

type EnemyDefendComponent struct {
	MaxHP       int
	HP          int
	DodgeChance float32
	Armour      int
	Named       // will just be a pointer to the method that implements this
	Criticaller
	StatMultiplier
}

func (d *EnemyDefendComponent) TakeDamage(nDamage int, bBlockable bool) (bool, string) {
	bHit := true
	var str string
	if rand.Float32() > d.DodgeChance {
		if bBlockable {
			d.HP -= nDamage - d.Armour
		} else {
			d.HP -= nDamage
		}
		str = fmt.Sprintf("%v now has %v/%v health\n",
			d.GetName(),
			d.HP,
			d.MaxHP)
	} else {
		fmt.Println("Attack dodged, no damage was dealt")
		bHit = false
	}
	return bHit, str
}

func (d EnemyDefendComponent) IsDead() bool {
	if d.HP <= 0 {
		return true
	} else {
		return false
	}
}

type EnemyAttackComponent struct {
	weapon Weapon
	Named
	Criticaller
	StatMultiplier
}

// TODO: Create a NewEnemy contructor that initializes each components interfaces

// This function will take in an enemy pointer, and initialize the components interfaces to
// the enemy itself, so that the components will have access to the methods they need.
// Call this after creating the Enemy instance with the values you want.
func initializeEnemyInterfaces(enemy *Enemy) {

	// set up the enemy instance to satisfy be assigned to its components interfaces
	enemy.EnemyAttackComponent.Named = enemy
	enemy.EnemyAttackComponent.Criticaller = enemy
	enemy.EnemyAttackComponent.StatMultiplier = enemy

	enemy.EnemyDefendComponent.Named = enemy
	enemy.EnemyDefendComponent.Criticaller = enemy
	enemy.EnemyDefendComponent.StatMultiplier = enemy
}

func NewBozo(weapon Weapon, level int) *Enemy {

	enemy := Enemy{
		CharacterComponent{CharacterName: "Bozo", CatchPhrases: []string{"I'll bozo you!", "Get smoked bozo"}},
		EnemyAttackComponent{weapon: weapon},
		EnemyDefendComponent{MaxHP: 5, HP: 5, DodgeChance: 0.05, Armour: 1},
		level,
	}

	initializeEnemyInterfaces(&enemy)

	return &enemy
}

func (a EnemyAttackComponent) Attack(d IDefend) {
	isCritical := a.isCritical()

	damageToDeal := a.calculateDamage(a.weapon, isCritical)
	bHit, str := d.TakeDamage(damageToDeal, true)
	if bHit {
		fmt.Printf("%v used their %v and dealt %v damage.\n"+str,
			a.GetName(),
			a.weapon.Name,
			damageToDeal)

		if isCritical {
			fmt.Println("Critical Hit!")
		}
	}

}

func (a EnemyAttackComponent) calculateDamage(weapon Weapon, isCritical bool) int {
	damage := int(float32(weapon.BaseDamage) * a.GetStatMultiplier())

	if isCritical {
		damage *= int(a.weapon.CriticalBonus)
	}

	return damage
}

func (a EnemyAttackComponent) isCritical() bool {

	rand := rand.Float32()
	// if the random float is less than crit chance, it is a critical hit
	return rand < a.GetCritChance()
}

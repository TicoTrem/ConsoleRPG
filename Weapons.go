package main

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

// return a pointer to avoid copying the entire object when this is called
func NewTrainingSword() *Weapon {
	return &Weapon{
		Name:             "Training Sword",
		BaseDamage:       10,
		WeaponDamageType: Regular,
		DescriptorWord:   "Dull",
		AttackMessage:    "It might have done something...",
		CritChance:       0.12,
		CriticalBonus:    1.50,
	}
}

func NewFists() *Weapon {
	return &Weapon{
		Name:             "Fists",
		BaseDamage:       5,
		WeaponDamageType: Regular,
		DescriptorWord:   "Weak",
		AttackMessage:    "They probably felt something...",
		CritChance:       0.10,
		CriticalBonus:    1.5,
	}
}

package main

type Weapon interface {
	CalculateDamage(attr *Attribute, dice *Dice) int
	GetRank() int
	GetName() string
}

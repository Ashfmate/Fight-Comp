package main

type Fist struct {
	name string
	rank int
}

func NewFist() Fist {
	return Fist{"fist", 0}
}

func (Fist) CalculateDamage(attr *Attribute, dice *Dice) int {
	return attr.power + Roll(dice, 2)
}

func (self Fist) GetRank() int {
	return self.rank
}

func (self Fist) GetName() string {
	return self.name
}

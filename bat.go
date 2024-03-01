package main

type Bat struct {
	name string
	rank int
}

func NewBat() Bat {
	return Bat{"bat", 1}
}

func (Bat) CalculateDamage(attr *Attribute, dice *Dice) int {
	return attr.power*2 + Roll(dice, 1)
}

func (self Bat) GetRank() int {
	return self.rank
}

func (self Bat) GetName() string {
	return self.name
}

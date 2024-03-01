package main

type Knive struct {
	name string
	rank int
}

func NewKnive() Knive {
	return Knive{"knive", 2}
}

func (Knive) CalculateDamage(attr *Attribute, dice *Dice) int {
	return attr.speed*3 + Roll(dice, 3)
}

func (self Knive) GetRank() int {
	return self.rank
}

func (self Knive) GetName() string {
	return self.name
}

package main

type Attribute struct {
	hp, speed, power int
}

func NewAttribute(hp, speed, power int) Attribute {
	return Attribute{hp, speed, power}
}

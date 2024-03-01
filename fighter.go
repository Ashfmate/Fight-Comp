package main

type Fighter interface {
	IsAlive() bool
	GetInitiative() int
	Attack(other *Fighter)
	Tick()
	SpecialMove(other *Fighter)
	GiveWeapon(wep *Weapon)
	PilferWeapon() *Weapon
	HasWeapon() bool
	roll(n_roll int) int
	applyDamageTo(other *Fighter, damage int)
	takeDamage(damage int)
	GetName() string
	GetWeapon() Weapon
	GetAttributes() Attribute
}

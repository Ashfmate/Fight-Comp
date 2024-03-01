package main

import "fmt"

type Warrior struct {
	wep  *Weapon
	dice Dice
	name string
	attr Attribute
}

func NewWarrior(name string, dice Dice, wep *Weapon) Fighter {
	fmt.Printf("%s has entered the ring!\n", name)
	var res Fighter
	res = &Warrior{
		name: name,
		dice: dice,
		attr: NewAttribute(80, 4, 10),
		wep:  wep,
	}
	return res
}

func (self *Warrior) roll(n_roll int) int {
	return Roll(&self.dice, n_roll)
}

func (self *Warrior) applyDamageTo(other *Fighter, damage int) {
	(*other).takeDamage(damage)
	target := *other
	fmt.Printf("%s takes %d damage\n", target.GetName(), damage)
	if !target.IsAlive() {
		fmt.Printf("%s has been killed by %s\n", target.GetName(), self.name)
	}
}

func (self *Warrior) IsAlive() bool {
	return self.attr.hp > 0
}

func (self *Warrior) GetInitiative() int {
	return self.attr.speed + self.roll(2)
}

func (self *Warrior) GetName() string {
	return self.name
}

func (self *Warrior) GetAttributes() Attribute {
	return self.attr
}

func (self *Warrior) GetWeapon() Weapon {
	return *self.wep
}

func (self *Warrior) takeDamage(damage int) {
	self.attr.hp -= damage
}

func (self *Warrior) Attack(other *Fighter) {
	if self.IsAlive() && (*other).IsAlive() {
		fmt.Printf("%s attacks %s with %s\n", self.name, (*other).GetName(), (*self.wep).GetName())
		self.applyDamageTo(other, (*self.wep).CalculateDamage(&self.attr, &self.dice))
	}
}

func (self *Warrior) Tick() {
	if self.IsAlive() {
		rec := self.roll(1)
		fmt.Printf("%s recovered %d\n", self.name, rec)
		self.attr.hp += rec
	}
}

func (self *Warrior) SpecialMove(other *Fighter) {
	if !self.IsAlive() || !(*other).IsAlive() {
		return
	}
	if self.roll(1) <= 3 {
		fmt.Printf("%s forgot his family\n", self.name)
		return
	}
	fmt.Printf("%s remembers his family and becomes super %s\n", self.name, self.GetName())
	self.name = "Super " + self.name
	self.attr.hp += 10
	self.attr.speed += 3
	self.attr.power = (self.attr.power * 60) / 40
}

func (self *Warrior) GiveWeapon(wep *Weapon) {
	self.wep = wep
}

func (self *Warrior) PilferWeapon() *Weapon {
	ret := self.wep
	self.wep = nil
	return ret
}

func (self *Warrior) HasWeapon() bool {
	return self.wep != nil
}

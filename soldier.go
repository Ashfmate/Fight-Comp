package main

import "fmt"

type Soldier struct {
	wep  *Weapon
	dice Dice
	name string
	attr Attribute
}

func NewSoldier(name string, dice Dice, wep *Weapon) Fighter {
	fmt.Printf("%s has entered the ring!\n", name)
	var res Fighter
	res = &Warrior{
		name: name,
		dice: dice,
		attr: NewAttribute(70, 7, 14),
		wep:  wep,
	}
	return res
}

func (self *Soldier) roll(n_roll int) int {
	return Roll(&self.dice, n_roll)
}

func (self *Soldier) applyDamageTo(other *Fighter, damage int) {
	(*other).takeDamage(damage)
	target := *other
	fmt.Printf("%s takes %d damage\n", target.GetName(), damage)
	if !target.IsAlive() {
		fmt.Printf("%s has been killed by %s\n", target.GetName(), self.name)
	}
}

func (self *Soldier) IsAlive() bool {
	return self.attr.hp > 0
}

func (self *Soldier) GetInitiative() int {
	return self.attr.speed + self.roll(2)
}

func (self *Soldier) GetName() string {
	return self.name
}

func (self *Soldier) GetAttributes() Attribute {
	return self.attr
}

func (self *Soldier) GetWeapon() Weapon {
	return *self.wep
}

func (self *Soldier) takeDamage(damage int) {
	fmt.Printf("%s's health is %d\n", self.name, self.attr.hp)
	self.attr.hp -= damage
	fmt.Printf("%s's health is now %d\n", self.name, self.attr.hp)
}

func (self *Soldier) Attack(other *Fighter) {
	if self.IsAlive() && (*other).IsAlive() {
		fmt.Printf("%s attacks %s with his %s\n", self.name, (*other).GetName(), (*self.wep).GetName())
		self.applyDamageTo(other, (*self.wep).CalculateDamage(&self.attr, &self.dice))
	}
}

func (self *Soldier) Tick() {
	if self.IsAlive() {
		fmt.Printf("%s is hurt by a stone\n", self.name)
		damage := self.roll(3) + 20
		self.takeDamage(damage)
		fmt.Printf("%s takes %d damage\n", self.name, damage)
		if !self.IsAlive() {
			fmt.Printf("%s killed himself\n", self.name)
			return
		}
		rec := self.roll(1)
		fmt.Printf("%s recovered %d\n", self.name, rec)
		self.attr.hp += rec
	}
}

func (self *Soldier) SpecialMove(other *Fighter) {
	if !self.IsAlive() || !(*other).IsAlive() {
		return
	}
	if self.roll(1) <= 4 {
		fmt.Printf("%s's gun is not responding\n", self.name)
		return
	}
	fmt.Printf("%s attacks %s with an AK-47\n", self.name, (*other).GetName())
	self.applyDamageTo(other, self.roll(3)+20)
}

func (self *Soldier) GiveWeapon(wep *Weapon) {
	self.wep = wep
}

func (self *Soldier) PilferWeapon() *Weapon {
	ret := self.wep
	self.wep = nil
	return ret
}

func (self *Soldier) HasWeapon() bool {
	return self.wep != nil
}

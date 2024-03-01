package main

import (
	"math/rand"
	"time"
)

type Dice struct {
	rng rand.Rand
}

func NewDice() Dice {
	return Dice{*rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func Roll(dice *Dice, nDice int) int {
	total := 0
	for range nDice {
		total += dice.rng.Intn(6) + 1
	}
	return total
}

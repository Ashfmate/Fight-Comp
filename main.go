package main

import (
	"fmt"
	"slices"
	"sort"
)

func readLine[T any]() (*T, error) {
	buffer := ""
	_, err := fmt.Scanln(&buffer)
	if err != nil {
		return nil, err
	}
	var res T
	_, err = fmt.Sscan(buffer, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func main() {
	fist := NewFist()
	bat := NewBat()
	knive := NewKnive()
	dice := NewDice()
	wep := []Weapon{fist, bat, knive}

	fmt.Println("++++++++++++++++++Team 1++++++++++++++++++")
	team1 := []Fighter{
		NewWarrior("Ahmed", dice, &wep[Roll(&dice, 1)%3]),
		NewSoldier("Brhoom", dice, &wep[Roll(&dice, 1)%3]),
		NewWarrior("Mohammed", dice, &wep[Roll(&dice, 1)%3]),
	}

	fmt.Println("++++++++++++++++++Team 1++++++++++++++++++")

	team2 := []Fighter{
		NewSoldier("Ouais", dice, &wep[Roll(&dice, 1)%3]),
		NewWarrior("Taha", dice, &wep[Roll(&dice, 1)%3]),
		NewSoldier("Noory", dice, &wep[Roll(&dice, 1)%3]),
	}

	fmt.Println("-----------------------------------BEGIN-----------------------------------")

	for team1Alive, team2Alive := getAlive(team1, team2); team1Alive && team2Alive; team1Alive, team2Alive = getAlive(team1, team2) {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++ Team 1 Status ++++++++++++++++++++++++++++++++++++++++++++")
		for _, elem := range team1 {
			if !elem.IsAlive() {
				continue
			}
			fmt.Println(elem.GetName(), elem.GetAttributes(), "and his weapon is", elem.GetWeapon().GetName())
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++ Team 2 Status ++++++++++++++++++++++++++++++++++++++++++++")
		for _, elem := range team2 {
			if !elem.IsAlive() {
				continue
			}
			fmt.Println(elem.GetName(), elem.GetAttributes(), "and his weapon is", elem.GetWeapon().GetName())
		}

		// Shuffling and partitioning team1
		dice.rng.Shuffle(len(team1), func(i, j int) { team1[i], team1[j] = team1[j], team1[i] })
		sort.Slice(team1, func(i, j int) bool { return team1[i].IsAlive() })
		// Shuffling and partitioning team2
		dice.rng.Shuffle(len(team1), func(i, j int) { team2[i], team2[j] = team2[j], team2[i] })
		sort.Slice(team2, func(i, j int) bool { return team2[i].IsAlive() })

		for i := range 3 {
			engage(&(team1[i]), &(team2[i]))
			doSpecials(&team1[i], &team2[i])
			fmt.Println("------------------------------------")
		}
		fmt.Println("=====================================")

		for i := range 3 {
			team1[i].Tick()
			team2[i].Tick()
		}

		fmt.Println("=====================================")
		fmt.Println("Press the enter key to continue")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++ Team 1 Status ++++++++++++++++++++++++++++++++++++++++++++")
		for _, elem := range team1 {
			if !elem.IsAlive() {
				continue
			}
			fmt.Println(elem.GetName(), elem.GetAttributes(), "and his weapon is", elem.GetWeapon().GetName())
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++ Team 2 Status ++++++++++++++++++++++++++++++++++++++++++++")
		for _, elem := range team2 {
			if !elem.IsAlive() {
				continue
			}
			fmt.Println(elem.GetName(), elem.GetAttributes(), "and his weapon is", elem.GetWeapon().GetName())
		}
		readLine[string]()
	}

	team1Alive, _ := getAlive(team1, team2)
	if team1Alive {
		fmt.Println("Team 1 is Victorious!!!")
	} else {
		fmt.Println("Team 2 is Victorious!!!")
	}
}

func getAlive(team1 []Fighter, team2 []Fighter) (bool, bool) {
	return slices.ContainsFunc(team1, func(elem Fighter) bool { return elem.IsAlive() }),
		slices.ContainsFunc(team2, func(elem Fighter) bool { return elem.IsAlive() })
}

func takeWeaponIfDead(taker *Fighter, giver *Fighter) {
	_taker, _giver := *taker, *giver
	if _taker.IsAlive() && !_giver.IsAlive() && _giver.HasWeapon() {
		if _giver.GetWeapon().GetRank() <= _taker.GetWeapon().GetRank() {
			return
		}
		fmt.Printf("%s takes from %s %s\n", _taker.GetName(), _giver.GetName(), _giver.GetWeapon().GetName())
		(*taker).GiveWeapon((*giver).PilferWeapon())
	}
}

func orderInit(f1 *Fighter, f2 *Fighter) (*Fighter, *Fighter) {
	if (*f1).GetInitiative() > (*f2).GetInitiative() {
		return f1, f2
	}
	return f2, f1
}

func engage(f1 *Fighter, f2 *Fighter) {
	loc_f1, loc_f2 := orderInit(f1, f2)
	(*loc_f1).Attack(loc_f2)
	takeWeaponIfDead(loc_f1, loc_f2)
	(*loc_f2).Attack(loc_f1)
	takeWeaponIfDead(loc_f2, loc_f1)
}

func doSpecials(f1 *Fighter, f2 *Fighter) {
	loc_f1, loc_f2 := orderInit(f1, f2)
	(*loc_f1).SpecialMove(loc_f2)
	takeWeaponIfDead(loc_f1, loc_f2)
	(*loc_f2).SpecialMove(loc_f1)
	takeWeaponIfDead(loc_f2, loc_f1)
}

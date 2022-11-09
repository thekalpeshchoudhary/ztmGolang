//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package rcv

import "fmt"

type Health struct {
	health    int
	maxHealth int
}

type Energy struct {
	energy    int
	maxEnergy int
}

type Player struct {
	name       string
	healthStat Health
	energyStat Energy
}

func (player *Player) ChangeHealth(howMuch int, action string) {
	fmt.Println("\n---Change Health by", howMuch, ", action", action)
	if action == "add" {
		if player.healthStat.health+howMuch > player.healthStat.maxHealth {
			player.healthStat.health = player.healthStat.maxHealth
			fmt.Println("Max Health can only be :", player.healthStat.maxHealth)
		} else {
			player.healthStat.health += howMuch
		}
	} else {
		if player.healthStat.health-howMuch < 0 {
			fmt.Println("Health cannot be negative")
		} else {
			player.healthStat.health -= howMuch
		}
	}
	fmt.Println("Player Name :", player.name)
	fmt.Println("Player Health :", player.healthStat.health)
	fmt.Println("Player Energy :", player.energyStat.energy)
}

func (player *Player) ChangeEnergy(howMuch int, action string) {
	fmt.Println("\n---Change Energy by", howMuch, ", action", action)
	if action == "add" {
		if player.energyStat.energy+howMuch > player.energyStat.maxEnergy {
			player.energyStat.energy = player.energyStat.maxEnergy
			fmt.Println("Max Energy can only be :", player.energyStat.maxEnergy)
		} else {
			player.energyStat.energy += howMuch
		}
	} else {
		if player.energyStat.energy-howMuch < 0 {
			fmt.Println("Energy cannot be negative")
		} else {
			player.energyStat.energy -= howMuch
		}
	}
	fmt.Println("Player Name :", player.name)
	fmt.Println("Player Health :", player.healthStat.health)
	fmt.Println("Player Energy :", player.energyStat.energy)
}

func main() {
	thekc66 := Player{
		name:       "theKC66",
		energyStat: Energy{energy: 100, maxEnergy: 200},
		healthStat: Health{health: 100, maxHealth: 100},
	}
	freeman := Player{
		name:       "Freeman",
		energyStat: Energy{energy: 100, maxEnergy: 200},
		healthStat: Health{health: 90, maxHealth: 100},
	}
	aarya := Player{
		name:       "theAC66",
		energyStat: Energy{energy: 100, maxEnergy: 200},
		healthStat: Health{health: 82, maxHealth: 100},
	}

	thekc66.ChangeEnergy(50, "add")
	thekc66.ChangeHealth(20, "sub")

	freeman.ChangeEnergy(100, "add")
	freeman.ChangeHealth(100, "sub")

	aarya.ChangeEnergy(120, "sub")
	aarya.ChangeHealth(18, "add")
}

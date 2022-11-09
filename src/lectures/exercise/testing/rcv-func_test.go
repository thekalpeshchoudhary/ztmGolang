// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package rcv

import "testing"

func TestChangeHealthEnergy(t *testing.T) {
	thekc66 := Player{
		name:       "theKC66",
		energyStat: Energy{energy: 100, maxEnergy: 200},
		healthStat: Health{health: 100, maxHealth: 100},
	}
	thekc66.ChangeEnergy(150, "add")
	if thekc66.energyStat.energy > thekc66.energyStat.maxEnergy {
		t.Errorf("Should not have gone above max energy")
	}
	thekc66.ChangeEnergy(300, "sub")
	if thekc66.energyStat.energy < 0 {
		t.Errorf("Should not have gone below 0")
	}

	thekc66.ChangeHealth(150, "add")
	if thekc66.healthStat.health > thekc66.healthStat.maxHealth {
		t.Errorf("Should not have gone above max Health")
	}
	thekc66.ChangeEnergy(300, "sub")
	if thekc66.healthStat.health < 0 {
		t.Errorf("Health Should not have gone below 0")
	}
}

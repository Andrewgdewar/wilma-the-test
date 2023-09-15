package challenges

import (
	"math/rand"
	"time"
)

// There are a number of things that are broken with the "UltraFight" below

// Task 1:
// For some reason, the fight isn't finishing and no winner is chosen?
// Help me figure out why and fix it for me will you!

// Task 2:
// The winning fighter is not receiving experience (increased powerlevel) after winning a fight.
// This is all wrong, in DragonBall when you win a fight your "Power Level" increases.
// Make it so that the winner of the fight gains 1 powerlevel
// Also, fighters should recover quickly!
// Reset their health to 100 after each fight has concluded.

// Task 3:
// Keeping the above in mind, make the fighters fight until one of them reaches a POWER LEVEL 9000!
// They are the new "Winning fighter", and should be returned by the function.

// Task 4 (Stretch)
// Wow, those fights sure took a while!
// How could you go about using them to speed up this fight?
// Are you familiar with Goroutines?

// Player type
type Player struct {
	Name       string
	Health     int
	PowerLevel int
}

func kamehamehaBeamAttack(player Player) int {
	explosionDamage := rand.Intn(50)
	player.Health -= explosionDamage
	return explosionDamage
}

// UltraFight ...
func UltraFight() Player {
	// Turn this on to see the fight!
	debug := false

	piccolo := Player{
		Name:       "Piccolo",
		Health:     100,
		PowerLevel: 0,
	}

	goku := Player{
		Name:       "Goku",
		Health:     100,
		PowerLevel: 0,
	}

	winningPlayer := Player{}
	count := 20

	//Fight
	for i := 1; i < count; i++ {
		// If even number Piccolo attacks
		if i%2 == 0 {
			hitAmount := kamehamehaBeamAttack(goku)
			if debug {
				print("Piccolo attacks Goku for ", hitAmount, " damage!\n")
			}
			if goku.Health <= 0 {
				winningPlayer = piccolo
				count = 0
			}
			// If odd number Goku attacks
		} else {
			hitAmount := kamehamehaBeamAttack(piccolo)
			if debug {
				print("Goku attacks Piccolo for ", hitAmount, " damage!\n")
			}
			if piccolo.Health <= 0 {
				winningPlayer = goku
				count = 0
			}
		}
		// This is to represent their super speed fighting time.
		time.Sleep(time.Microsecond * 50)
	}

	return winningPlayer
}

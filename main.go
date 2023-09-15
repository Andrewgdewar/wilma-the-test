package main

import (
	"fmt"
	"reflect"
	"time"
	"wilma-the-test/challenges"
)

func main() {
	print("\nTwo Sum!\n")
	// TwoSum
	if true {

		start := time.Now()
		numberOfTimes := 5000

		nums := [][]int{
			{2, 7, 11, 15},
			{3, 2, 4},
			{3, 3},
		}

		targets := []int{
			9,
			6,
			6,
		}

		results := [][]int{
			{0, 1},
			{1, 2},
			{0, 1},
		}

		for outerIdx := 0; outerIdx <= numberOfTimes; outerIdx++ {
			for idx := range nums {
				result := challenges.TwoSum(nums[idx], targets[idx])
				if outerIdx == numberOfTimes && reflect.DeepEqual(result, results[idx]) {
					fmt.Print(idx+1, " ✅ 🙌 🏆 🥇 🎯 💯\n")
				} else if outerIdx == numberOfTimes {
					fmt.Print(idx+1, " ❌")
					if len(result) > 1 {
						fmt.Print(result[0], result[1])
					}
					print("\n")
				}
			}
		}
		elapsed := time.Since(start)
		fmt.Printf("\nTook %s", elapsed)
	}

	//Ultra fight!
	if true {
		print("\n\nUltra Fight!\n")
		start := time.Now()
		winningPlayer := challenges.UltraFight()

		if winningPlayer.Name == "" {
			fmt.Print("❌ The fight is broken!\n")
		} else {
			fmt.Print("✅ A player won! Wahoo!\n")
		}

		if winningPlayer.PowerLevel == 0 {
			fmt.Print("❌ No experience gained!\n")
		} else {
			fmt.Print("✅ Experience granted!\n", winningPlayer.PowerLevel)
		}

		if winningPlayer.PowerLevel != 9000 {
			fmt.Print("❌ Power level is unimpressive!\n")
		} else {
			fmt.Print("✅ Woohooo almost super saiyan!\n", winningPlayer.PowerLevel)
		}

		print("\n---Winning player---\nName: ", winningPlayer.Name,
			"\nHealth: ", winningPlayer.Health,
			"\nPowerLevel: ", winningPlayer.PowerLevel, "\n")
		fmt.Printf("\nTook %s\n", time.Since(start))
	}

}

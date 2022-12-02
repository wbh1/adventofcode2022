package solutions

import (
	"fmt"
	"strings"
)

type DayTwo struct {
	input []string
}

const (
	win  = 6
	draw = 3
	loss = 0
)

var TheirPlay = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

// PartOne
var MyPlay = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

// PartTwo
var MyPlayOutcome = map[string]string{
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

var OutcomeMatrix = map[string]map[string]string{
	"Rock": {
		"Rock":     "Draw",
		"Paper":    "Lose",
		"Scissors": "Win",
	},

	"Paper": {
		"Rock":     "Win",
		"Paper":    "Draw",
		"Scissors": "Lose",
	},

	"Scissors": {
		"Rock":     "Lose",
		"Paper":    "Win",
		"Scissors": "Draw",
	},
}

var ActionValue = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

func (d *DayTwo) SetInput(input []string) {
	d.input = input
}

func (d *DayTwo) PartOne() string {
	score := 0

	for _, round := range d.input {
		actions := strings.Split(round, " ")
		if len(actions) != 2 {
			panic("Missing two actions for round.")
		}

		tp := TheirPlay[actions[0]]
		mp := MyPlay[actions[1]]

		// Add my choice to the score
		score += ActionValue[mp]
		switch OutcomeMatrix[mp][tp] {
		case "Win":
			score += win
		case "Lose":
			score += loss
		case "Draw":
			score += draw
		}

	}

	return fmt.Sprintf("%d", score)
}

func (d *DayTwo) PartTwo() string {
	score := 0

	for _, round := range d.input {
		actions := strings.Split(round, " ")
		if len(actions) != 2 {
			panic("Missing two actions for round.")
		}

		tp := TheirPlay[actions[0]]
		myOutcome := MyPlayOutcome[actions[1]]

		for mp, theirOutcome := range OutcomeMatrix[tp] {
			if myOutcome == "Draw" && theirOutcome == "Draw" {
				score += ActionValue[mp] + draw
				break
			} else if myOutcome == "Win" && theirOutcome == "Lose" {
				score += ActionValue[mp] + win
				break
			} else if myOutcome == "Lose" && theirOutcome == "Win" {
				score += ActionValue[mp] + loss
				break
			}
		}

	}
	return fmt.Sprintf("%d", score)
}

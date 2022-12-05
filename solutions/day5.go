package solutions

import (
	"regexp"
	"unicode"

	"github.com/wbh1/adventofcode2022/util"
)

const CrateMovementRegex = `move (?P<Qty>\d+) from (?P<Src>\d+) to (?P<Dst>\d+)`

type CrateMovement struct {
	Qty, Src, Dst int
}

type DayFive struct {
	input        []string
	CrateStacks  []util.Stack
	Instructions []CrateMovement
}

func (d *DayFive) SetInput(input []string) {
	d.input = input
	instructionsStartIndex := 0

	// Need to provision the number of stacks first
	for i, line := range input {
		if line == "" {
			numStacks := util.MustStringToInt(string(input[i-1][len(input[i-1])-2]))
			d.CrateStacks = make([]util.Stack, numStacks)
			instructionsStartIndex = i + 1
			break
		}
	}

	// Initial setup of crates
	for i := instructionsStartIndex - 2; i >= 0; i-- {
		line := input[i]
		// Start at index 1 of string (2nd character) and every 4 characters will have the next letter
		for i := 1; i < len(line); i += 4 {
			if unicode.IsSpace(rune(line[i])) {
				continue
			}
			// This is the line with the stack numbers, skip it.
			if unicode.IsDigit(rune(line[i])) {
				break
			}
			d.CrateStacks[i/4].Push(string(line[i]))
		}
	}

	// Prevent adding instructions again if setup is called twice
	if d.Instructions != nil {
		return
	}
	cr := regexp.MustCompile(CrateMovementRegex)
	for _, line := range input[instructionsStartIndex:] {
		matches := cr.FindStringSubmatch(line)
		d.Instructions = append(d.Instructions, CrateMovement{
			Qty: util.MustStringToInt(matches[cr.SubexpIndex("Qty")]),
			Src: util.MustStringToInt(matches[cr.SubexpIndex("Src")]) - 1,
			Dst: util.MustStringToInt(matches[cr.SubexpIndex("Dst")]) - 1,
		})
	}
}

func (d *DayFive) topCrates() string {
	topCrates := ""
	for _, crate := range d.CrateStacks {
		topCrates += crate.Top.Value
	}
	return topCrates
}

func (d *DayFive) PartOne() string {
	for _, mvmt := range d.Instructions {
		for i := 1; i <= mvmt.Qty; i++ {
			add := d.CrateStacks[mvmt.Src].Pop()
			if add != nil {
				d.CrateStacks[mvmt.Dst].Push(add.Value)
			} else {
				panic("Somehow we're trying to move a crate that doesn't exist!")
			}
		}
	}

	return d.topCrates()
}

func (d *DayFive) PartTwo() string {
	// Must reset input to undo movement from Part1
	d.SetInput(d.input)
	for _, mvmt := range d.Instructions {
		// Create a new stack of additions to ensure they get added back in the right order
		additions := util.Stack{}
		for i := 1; i <= mvmt.Qty; i++ {
			additions.Push(d.CrateStacks[mvmt.Src].Pop().Value)
		}
		for i := 1; i <= mvmt.Qty; i++ {
			d.CrateStacks[mvmt.Dst].Push(additions.Pop().Value)
		}
	}

	return d.topCrates()
}

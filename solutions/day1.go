package solutions

import (
	"fmt"
	"sort"
	"strconv"
)

type DayOne struct {
	input []string
}

var (
	elves []Elf
)

type Elf struct {
	Food          []int
	TotalCalories int
}

func (d *DayOne) SetInput(input []string) {
	d.input = input

}

func (d *DayOne) PartOne() string {
	elf := Elf{}

	for _, line := range d.input {
		// Newline
		if line == "" {
			elves = append(elves, elf)
			elf = Elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		elf.Food = append(elf.Food, calories)
		elf.TotalCalories += calories
	}

	max_cals := 0
	for _, elf := range elves {
		if elf.TotalCalories > max_cals {
			max_cals = elf.TotalCalories
		}
	}

	return fmt.Sprintf("%d", max_cals)
}

func (d *DayOne) PartTwo() string {
	sortElves := func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	}
	sort.Slice(elves, sortElves)

	return fmt.Sprintf("%d", elves[0].TotalCalories+elves[1].TotalCalories+elves[2].TotalCalories)
}

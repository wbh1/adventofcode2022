package solutions

import (
	"fmt"
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
	// 203731 too low
	top3 := [3]Elf{}
	for _, elf := range elves {
		switch cals := elf.TotalCalories; {
		case cals > top3[0].TotalCalories:
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = elf
		case cals > top3[1].TotalCalories:
			top3[2] = top3[1]
			top3[1] = elf
		case cals > top3[2].TotalCalories:
			top3[2] = elf
		}
	}

	total_cals := 0
	for _, elf := range top3 {
		total_cals += elf.TotalCalories
	}

	return fmt.Sprintf("%d", total_cals)
}

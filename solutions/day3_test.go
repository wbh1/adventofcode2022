package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayThree = DayTest{
		Input: strings.Split(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, "\n"),
		AlwaysRunPartOne: true,
		Day:              &solutions.DayThree{},
		Part1Result:      "157",
		Part2Result:      "70",
	}
)

func TestDayThreePartOne(t *testing.T) {
	DayThree.TestDayPartOne(t)
}

func TestDayThreePartTwo(t *testing.T) {
	DayThree.TestDayPartTwo(t)
}

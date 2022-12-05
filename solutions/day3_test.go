package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	testInputDay3 = strings.Split(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, "\n")

	day3 = solutions.DayThree{}
)

func TestDayThreePartOne(t *testing.T) {
	day3.SetInput(testInputDay3)

	p1 := day3.PartOne()
	if p1 != "157" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "157", p1)
	}
}

func TestDayThreePartTwo(t *testing.T) {
	day3.SetInput(testInputDay3)

	day3.PartOne()
	p2 := day3.PartTwo()
	if p2 != "70" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "70", p2)
	}
}

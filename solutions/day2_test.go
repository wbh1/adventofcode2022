package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	testInput = strings.Split(`A Y
B X
C Z`, "\n")

	day = solutions.DayTwo{}
)

func setup(t *testing.T) {
	day.SetInput(testInput)
}

func TestDayTwoPartOne(t *testing.T) {
	setup(t)

	p1 := day.PartOne()
	if p1 != "15" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "15", p1)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	setup(t)

	p2 := day.PartTwo()
	if p2 != "12" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "12", p2)
	}
}

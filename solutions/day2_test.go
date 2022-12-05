package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
	"github.com/wbh1/adventofcode2022/util"
)

var (
	dayTwoTestInput = strings.Split(`A Y
B X
C Z`, "\n")

	day = solutions.DayTwo{}
)

func setup(t *testing.T, input []string, day util.Solution) {
	day.SetInput(input)
}

func TestDayTwoPartOne(t *testing.T) {
	setup(t, dayTwoTestInput, &day)

	p1 := day.PartOne()
	if p1 != "15" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "15", p1)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	setup(t, dayTwoTestInput, &day)

	p2 := day.PartTwo()
	if p2 != "12" {
		t.Fatalf("Unexpected return. Want %s, Got %s", "12", p2)
	}
}

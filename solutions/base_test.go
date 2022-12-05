package solutions_test

import (
	"testing"

	"github.com/wbh1/adventofcode2022/util"
)

type DayTest struct {
	Input            []string
	AlwaysRunPartOne bool
	Day              util.Solution
	Part1Result      string
	Part2Result      string
}

func (dt *DayTest) setup(t *testing.T) {
	dt.Day.SetInput(dt.Input)
}

func (dt *DayTest) TestDayPartOne(t *testing.T) {
	dt.setup(t)

	p1 := dt.Day.PartOne()
	if p1 != dt.Part1Result {
		t.Fatalf("Unexpected return. Want %s, Got %s", dt.Part1Result, p1)
	}
}

func (dt *DayTest) TestDayPartTwo(t *testing.T) {
	dt.setup(t)

	if dt.AlwaysRunPartOne {
		dt.Day.PartOne()
	}
	p2 := dt.Day.PartTwo()
	if p2 != dt.Part2Result {
		t.Fatalf("Unexpected return. Want %s, Got %s", dt.Part2Result, p2)
	}
}

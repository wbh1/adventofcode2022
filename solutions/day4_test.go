package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayFour = DayTest{
		Input: strings.Split(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`, "\n"),
		AlwaysRunPartOne: true,
		Day:              &solutions.DayFour{},
		Part1Result:      "2",
		Part2Result:      "4",
	}
)

func TestDayFourPartOne(t *testing.T) {
	DayFour.TestDayPartOne(t)
}

func TestDayFourPartTwo(t *testing.T) {
	DayFour.TestDayPartTwo(t)
}

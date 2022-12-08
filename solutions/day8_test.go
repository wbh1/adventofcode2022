package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayEight = DayTest{Input: strings.Split(`30373
25512
65332
33549
35390`, "\n"),
		AlwaysRunPartOne: true,
		Day:              &solutions.DayEight{},
		Part1Result:      "21",
		Part2Result:      "8",
	}
)

func TestDayEightPartOne(t *testing.T) {
	DayEight.TestDayPartOne(t)
}

func TestDayEightPartTwo(t *testing.T) {
	DayEight.TestDayPartTwo(t)
}

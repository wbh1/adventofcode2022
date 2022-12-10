package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayNine = DayTest{Input: strings.Split(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`, "\n"),
		AlwaysRunPartOne: false,
		Day:              &solutions.DayNine{},
		Part1Result:      "13",
		Part2Result:      "1",
	}
)

func TestDayNinePartOne(t *testing.T) {
	DayNine.TestDayPartOne(t)
}

func TestDayNinePartTwo(t *testing.T) {
	DayNine.TestDayPartTwo(t)

	DayNine.Input = strings.Split(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`, "\n")

	DayNine.Part2Result = "36"

	DayNine.TestDayPartTwo(t)
}

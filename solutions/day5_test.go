package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayFive = DayTest{
		Input: strings.Split(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`, "\n"),
		AlwaysRunPartOne: false,
		Day:              &solutions.DayFive{},
		Part1Result:      "CMZ",
		Part2Result:      "MCD",
	}
)

func TestDayFivePartOne(t *testing.T) {
	DayFive.TestDayPartOne(t)
}

func TestDayFivePartTwo(t *testing.T) {
	DayFive.TestDayPartTwo(t)
}

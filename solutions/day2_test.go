package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DayTwo = DayTest{
		Input: strings.Split(`A Y
B X
C Z`, "\n"),
		Day:         &solutions.DayTwo{},
		Part1Result: "15",
		Part2Result: "12",
	}
)

func TestDayTwoPartOne(t *testing.T) {
	DayTwo.TestDayPartOne(t)
}

func TestDayTwoPartTwo(t *testing.T) {
	DayTwo.TestDayPartTwo(t)
}

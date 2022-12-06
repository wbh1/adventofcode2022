package solutions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DaySix = DayTest{Input: strings.Split(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`, "\n"),
		AlwaysRunPartOne: true,
		Day:              &solutions.DaySix{},
		Part1Result:      "7",
		Part2Result:      "19",
	}
)

func TestDaySixAllUnique(t *testing.T) {
	assert.True(t, solutions.AllUnique("abcd"))
	assert.False(t, solutions.AllUnique("adcd"))
}

func TestDaySixPartOne(t *testing.T) {
	DaySix.TestDayPartOne(t)

	additionalInputs := map[string]string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}

	for input, result := range additionalInputs {
		DaySix.Input = strings.Split(input, "\n")
		DaySix.Part1Result = result
		DaySix.TestDayPartOne(t)
	}
}

func TestDaySixPartTwo(t *testing.T) {
	additionalInputs := map[string]string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    "19",
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "23",
		"nppdvjthqldpwncqszvftbrmjlhg":      "23",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "29",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "26",
	}

	for input, result := range additionalInputs {
		DaySix.Input = strings.Split(input, "\n")
		DaySix.Part2Result = result
		DaySix.TestDayPartTwo(t)
	}
}

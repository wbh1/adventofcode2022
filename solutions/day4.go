package solutions

import (
	"fmt"
	"strings"

	"github.com/wbh1/adventofcode2022/util"
)

var overlaps int

type DayFour struct {
	input []string
}

func (d *DayFour) SetInput(input []string) {
	d.input = input
}

type ElfCleaner struct {
	Start int
	End   int
}

func (e ElfCleaner) Overlaps(z ElfCleaner) bool {
	// If any section between e's start and end is gte z's start and lte z's end, it overlaps
	for i := e.Start; i <= e.End; i++ {
		if z.Start <= i && i <= z.End {
			return true
		}
	}
	return false
}

func (e ElfCleaner) Contains(z ElfCleaner) bool {
	if e.Start <= z.Start && e.End >= z.End {
		return true
	}
	return false
}

func (d *DayFour) PartOne() string {
	result := 0
	for _, pair := range d.input {
		split := strings.Split(pair, ",")
		if len(split) != 2 {
			return fmt.Sprintf("ERROR: Didn't find two section ranges in %s", pair)
		}

		_elfA := strings.Split(split[0], "-")
		if len(_elfA) != 2 {
			return fmt.Sprintf("ERROR: Didn't find two numbers in %s (from %s)", split[0], pair)
		}
		ElfA := ElfCleaner{
			Start: util.MustStringToInt(_elfA[0]),
			End:   util.MustStringToInt(_elfA[1]),
		}

		_elfB := strings.Split(split[1], "-")
		if len(_elfB) != 2 {
			return fmt.Sprintf("ERROR: Didn't find two numbers in %s (from %s)", split[0], pair)
		}
		ElfB := ElfCleaner{
			Start: util.MustStringToInt(_elfB[0]),
			End:   util.MustStringToInt(_elfB[1]),
		}

		if ElfA.Contains(ElfB) || ElfB.Contains(ElfA) {
			result++
		}

		if ElfA.Overlaps(ElfB) {
			// fmt.Printf("%s overlaps with %s\n", split[0], split[1])
			overlaps++
		}

	}
	return fmt.Sprintf("%d", result)
}

func (d *DayFour) PartTwo() string {
	return fmt.Sprintf("%d", overlaps)
}

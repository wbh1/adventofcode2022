package solutions_test

import (
	"strings"
	"testing"

	"github.com/wbh1/adventofcode2022/solutions"
)

var (
	DaySeven = DayTest{Input: strings.Split(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`, "\n"),
		AlwaysRunPartOne: true,
		Day:              &solutions.DaySeven{},
		Part1Result:      "95437",
		Part2Result:      "24933642",
	}
)

func TestDaySevenPartOne(t *testing.T) {
	DaySeven.TestDayPartOne(t)
}

func TestDaySevenPartTwo(t *testing.T) {
	DaySeven.TestDayPartTwo(t)
}

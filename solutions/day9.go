package solutions

import (
	"fmt"
	"math"
	"strings"

	"github.com/wbh1/adventofcode2022/util"
)

type DayNine struct {
	input []string
}

type GridPoint struct {
	x int
	y int
}

type HeadPoint struct {
	loc               *GridPoint
	tail              *GridPoint
	TailPointsVisited util.Set
}

func distance(p1 *GridPoint, p2 *GridPoint) (x, y int) {
	x = int(math.Abs(float64(p1.x) - float64(p2.x)))
	y = int(math.Abs(float64(p1.y) - float64(p2.y)))
	return
}

func adjacent(p1 *GridPoint, p2 *GridPoint) bool {
	x, y := distance(p1, p2)

	return x <= 1 && y <= 1
}

func (hp *HeadPoint) Move(direction string, count int) {
	for i := 0; i < count; i++ {
		headStart := *hp.loc
		switch direction {
		case "U":
			hp.loc.y += 1
		case "D":
			hp.loc.y -= 1
		case "L":
			hp.loc.x -= 1
		case "R":
			hp.loc.x += 1
		}
		// If the head is no longer adjacent to the tail,
		// the tail moves to where the head just was.
		if !adjacent(hp.loc, hp.tail) {
			hp.tail.x = headStart.x
			hp.tail.y = headStart.y
			hp.TailPointsVisited.Add(fmt.Sprintf("(%d,%d)", hp.tail.x, hp.tail.y))
		}
	}
}

func (d *DayNine) SetInput(input []string) {
	d.input = input
}

func (d *DayNine) PartOne() string {
	head := HeadPoint{
		loc:               &GridPoint{},
		tail:              &GridPoint{},
		TailPointsVisited: util.Set{"(0,0)"},
	}

	for _, line := range d.input {
		inst := strings.Split(line, " ")
		dir := inst[0]
		count := util.MustStringToInt(inst[1])
		head.Move(dir, count)
	}
	return fmt.Sprintf("%d", len(head.TailPointsVisited))
}

func (d *DayNine) PartTwo() string {
	return "Not implemented."
}

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
	Parent            *HeadPoint
	loc               *GridPoint
	tails             []*HeadPoint
	TailPointsVisited util.Set
}

func distance(p1 *GridPoint, p2 *GridPoint) (x, y int) {
	x = p1.x - p2.x
	y = p1.y - p2.y
	return
}

func adjacent(p1 *GridPoint, p2 *GridPoint) bool {
	x, y := distance(p1, p2)

	return math.Abs(float64(x)) <= 1 && math.Abs(float64(y)) <= 1
}

func offset(n int) int {
	switch {
	case n == 0:
		return 0
	case n > 0:
		return 1
	default:
		return -1
	}
}

func (hp *HeadPoint) Move(direction string, count int) {
	for i := 0; i < count; i++ {
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

		for _, t := range hp.tails {
			if !adjacent(t.Parent.loc, t.loc) {
				x, y := distance(t.Parent.loc, t.loc)
				t.loc.x += offset(x)
				t.loc.y += offset(y)
				t.Parent.TailPointsVisited.Add(fmt.Sprintf("(%d,%d)", t.loc.x, t.loc.y))

			} else {
				break
			}
		}
	}
}

func (d *DayNine) SetInput(input []string) {
	d.input = input
}

func NewHeadPoint() *HeadPoint {
	return &HeadPoint{
		loc:               &GridPoint{},
		tails:             make([]*HeadPoint, 0),
		TailPointsVisited: util.Set{"(0,0)"},
	}
}

func (d *DayNine) PartOne() string {
	head := NewHeadPoint()
	head.tails = append(head.tails, NewHeadPoint())
	head.tails[0].Parent = head

	for _, line := range d.input {
		inst := strings.Split(line, " ")
		dir := inst[0]
		count := util.MustStringToInt(inst[1])
		head.Move(dir, count)
	}
	return fmt.Sprintf("%d", len(head.TailPointsVisited))
}

func (d *DayNine) PartTwo() string {
	head := NewHeadPoint()
	head.tails = make([]*HeadPoint, 9)
	for i := range head.tails {
		head.tails[i] = NewHeadPoint()
		if i < len(head.tails)-1 {
			head.tails[i].tails = head.tails[i+1:]
		}
		if i > 0 {
			head.tails[i].Parent = head.tails[i-1]
		} else {
			head.tails[i].Parent = head
		}
	}
	for _, line := range d.input {
		inst := strings.Split(line, " ")
		dir := inst[0]
		count := util.MustStringToInt(inst[1])
		head.Move(dir, count)
	}

	return fmt.Sprintf("%d", len(head.tails[7].TailPointsVisited))
}

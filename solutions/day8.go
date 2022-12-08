package solutions

import (
	"fmt"

	"github.com/wbh1/adventofcode2022/util"
)

type DayEight struct {
	input []string
}

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

var Trees [][]*Tree

func (d *DayEight) SetInput(input []string) {
	d.input = input
}

func (d *DayEight) PartOne() string {
	rows := len(d.input)
	cols := len(d.input[0])

	Trees = make([][]*Tree, rows)

	// Populate 2d array of trees
	for ri, row := range d.input {
		for ci, t := range row {
			tree := &Tree{
				Height: util.MustStringToInt(string(t)),
			}
			if ri == 0 || ri == rows-1 {
				tree.Visible = true
			} else if ci == 0 || ci == cols-1 {
				tree.Visible = true
			}
			Trees[ri] = append(Trees[ri], tree)
		}
	}

	// Determine which interior trees are visible
	// and what their scenic scores are
	for ri, row := range Trees {
		for ci, t := range row {
			// We already marked the edges as visible, so skip them.
			// Their scenic scores will always include a 0 so don't bother calculating them.
			if t.Visible {
				continue
			}

			rightVisible, leftVisible, upVisible, downVisible := true, true, true, true
			var rightScenic, leftScenic, upScenic, downScenic int

			// rightVisible
			if ci <= cols-2 {
				for _, rowNeighbor := range row[ci+1:] {
					rightScenic++
					if rowNeighbor.Height >= t.Height {
						rightVisible = false
						break
					}
				}
			}

			// leftVisible
			for i := ci - 1; i >= 0; i-- {
				leftScenic++
				if row[i].Height >= t.Height {
					leftVisible = false
					break
				}
			}

			// downVisible
			if ri <= rows-2 {
				for _, row := range Trees[ri+1:] {
					downScenic++
					if row[ci].Height >= t.Height {
						downVisible = false
						break
					}
				}
			}

			// upVisible
			for i := ri - 1; i >= 0; i-- {
				upScenic++
				if Trees[i][ci].Height >= t.Height {
					upVisible = false
					break
				}
			}

			t.Visible = upVisible || downVisible || leftVisible || rightVisible
			t.ScenicScore = upScenic * downScenic * leftScenic * rightScenic
		}
	}

	result := 0
	for _, row := range Trees {
		for _, tree := range row {
			if tree.Visible {
				result++
			}
		}
	}
	return fmt.Sprintf("%d", result)
}

func (d *DayEight) PartTwo() string {
	// Must've run PartOne first
	maxScenic := 0
	for _, row := range Trees {
		for _, tree := range row {
			if tree.ScenicScore > maxScenic {
				maxScenic = tree.ScenicScore
			}
		}
	}

	return fmt.Sprintf("%d", maxScenic)
}

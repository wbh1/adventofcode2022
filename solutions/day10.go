package solutions

import (
	"fmt"
	"math"
	"strings"

	"github.com/wbh1/adventofcode2022/util"
)

type DayTen struct {
	input []string
}

type Computer struct {
	Register     int
	Cycles       int
	Observations []int
	Display      [6][40]rune
	DisplayX     int
	DisplayY     int
}

var (
	c = Computer{
		Register: 1,
	}
)

func (c *Computer) drawPixel() rune {
	if math.Abs(float64(c.DisplayX-c.Register)) <= 1 {
		return '#'
	} else {
		return '.'
	}
}

func (c *Computer) RunCycle() {
	c.Display[c.DisplayY][c.DisplayX] = c.drawPixel()
	c.Cycles++
	if c.Cycles%40 == 20 {
		c.Observations = append(c.Observations, c.Cycles*c.Register)
		c.DisplayX++
	} else if c.Cycles%40 == 0 {
		c.DisplayX = 0
		c.DisplayY++
	} else {
		c.DisplayX++
	}

}

func (d *DayTen) SetInput(input []string) {
	d.input = input
}

func (d *DayTen) PartOne() string {
	for _, line := range d.input {
		inst := strings.Split(line, " ")
		// Noop
		if len(inst) == 1 {
			c.RunCycle()
			continue
		}

		if inst[0] == "addx" {
			for i := 0; i < 2; i++ {
				c.RunCycle()
			}
			c.Register += util.MustStringToInt(inst[1])
		} else {
			panic("Unknown instruction")
		}
	}

	result := 0
	for _, n := range c.Observations {
		result += n
	}

	return fmt.Sprintf("%d", result)
}

func (d *DayTen) PartTwo() string {
	for _, row := range c.Display {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Print("\n")
	}
	return "Look up."
}

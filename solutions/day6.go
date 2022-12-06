package solutions

import (
	"fmt"
	"strings"
)

type DaySix struct {
	input []string
}

func (d *DaySix) SetInput(input []string) {
	d.input = input
}

func AllUnique(s string) bool {
	for _, ltr := range s {
		if strings.Count(s, string(ltr)) > 1 {
			return false
		}
	}
	return true
}

func uniqueBufferEndIndex(s string, bufSize int) int {
	packet := s[:bufSize]

	if AllUnique(packet) {
		return bufSize
	}

	for i, ltr := range s {
		packet = packet[1:] + string(ltr)
		if AllUnique(packet) {
			return i + 1
		}
	}

	panic("No unique buffer found")
}

func (d *DaySix) PartOne() string {
	// Only 1 entry in input
	return fmt.Sprintf("%d", uniqueBufferEndIndex(d.input[0], 4))
}

func (d *DaySix) PartTwo() string {
	return fmt.Sprintf("%d", uniqueBufferEndIndex(d.input[0], 14))
}

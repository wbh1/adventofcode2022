package solutions

import (
	"fmt"
	"strings"

	"github.com/wbh1/adventofcode2022/util"
)

type DayThree struct {
	input []string
}

type Counter map[interface{}]int

type Rucksack struct {
	Contents          string
	CompartmentA      Counter
	CompartmentB      Counter
	ErrorItem         rune
	ErrorItemPriority int
}

var rucksacks []Rucksack

func (d *DayThree) SetInput(input []string) {
	d.input = input
}

func (d *DayThree) PartOne() string {
	result := 0
	for _, sack := range d.input {
		compartmentSize := len(sack) / 2
		r := Rucksack{Contents: sack}
		r.CompartmentA = make(Counter, compartmentSize)
		r.CompartmentB = make(Counter, compartmentSize)
		for _, item := range sack[:compartmentSize] {
			_, ok := r.CompartmentA[item]
			if ok {
				r.CompartmentA[item]++
			} else {
				r.CompartmentA[item] = 1
			}
		}
		for _, item := range sack[compartmentSize:] {
			_, ok := r.CompartmentB[item]
			if ok {
				r.CompartmentB[item]++
			} else {
				r.CompartmentB[item] = 1
			}
		}

		for item := range r.CompartmentA {
			if _, ok := r.CompartmentB[item]; ok {
				r.ErrorItem = item.(rune)
			}
		}

		if r.ErrorItem == 0 {
			return fmt.Sprintf("ERROR: No common items in %v and %v", r.CompartmentA, r.CompartmentB)
		}

		i := strings.IndexRune(util.AsciiLetters, r.ErrorItem)
		if i == -1 {
			return fmt.Sprintf("Uh... your error item isn't a letter? It's '%s'", string(r.ErrorItem))
		} else {
			r.ErrorItemPriority = i + 1
			result = result + i + 1
		}

		rucksacks = append(rucksacks, r)

	}
	return fmt.Sprintf("%d", result)
}

func (d *DayThree) PartTwo() string {
	groupSum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for _, item := range rucksacks[i].Contents {
			if strings.ContainsRune(rucksacks[i+1].Contents, item) && strings.ContainsRune(rucksacks[i+2].Contents, item) {
				groupSum += strings.IndexRune(util.AsciiLetters, item) + 1
				break
			}
		}
	}
	return fmt.Sprintf("%d", groupSum)
}

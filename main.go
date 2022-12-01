package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wbh1/adventofcode2022/solutions"
)

type Solution interface {
	PartOne() string
	PartTwo() string
}

var solved = map[string]Solution{
	"1": &solutions.DayOne{},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Must provide a day!\nUSAGE: <program> <day>\n")
		os.Exit(1)
	}

	day, ok := solved[os.Args[1]]
	if !ok {
		log.Fatalln("Undefined/unsolved day")
	}

	fmt.Println("Part 1:", day.PartOne())
	fmt.Println("Part 2:", day.PartTwo())
}

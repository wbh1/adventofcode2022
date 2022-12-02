package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wbh1/adventofcode2022/solutions"
	"github.com/wbh1/adventofcode2022/util"
)

type Solution interface {
	PartOne() string
	PartTwo() string
	SetInput([]string)
}

var solved = map[string]Solution{
	"1": &solutions.DayOne{},
	"2": &solutions.DayTwo{},
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

	input, err := util.ParseDayInput(os.Args[1])
	if err != nil {
		panic(err)
	}

	day.SetInput(input)

	fmt.Println("Part 1:", day.PartOne())
	fmt.Println("Part 2:", day.PartTwo())
}

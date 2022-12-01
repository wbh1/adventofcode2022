package util

import (
	"bufio"
	"fmt"
	"os"
)

func ParseDayInput(day uint8) ([]string, error) {
	var input = []string{}

	file, err := os.Open(fmt.Sprintf("./inputs/day%d.txt", day))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	} else {
		return input, nil
	}
}

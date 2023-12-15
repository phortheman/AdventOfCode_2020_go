package main

import (
	"fmt"
	"log"
	"strconv"

	day01 "github.com/phortheman/AdventOfCode2020_go/internal/Day_01"
	"github.com/phortheman/AdventOfCode2020_go/internal/file"
)

func main() {
	RunDay01()
}

func RunDay01() {
	var input []int

	filePath := "input/input_01.txt"
	lines, err := file.Read_Line(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Parsed line is not an number: ", err)
		}
		input = append(input, value)
	}

	part1, part2, err := day01.Solver(input)
	if err != nil {
		log.Fatal("Error with Day 01: ", err)
	}
	fmt.Printf("The answer for part 1: %v\nThe answer for part 2: %v\n", part1, part2)
}

package day01

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	var input []int = []int{979, 1721, 366, 299, 675, 1456} // Edited order to satify unit test coverage
	var resultPart1 int
	var foundPart1 bool
	for i := range input {
		if !foundPart1 {
			resultPart1, foundPart1 = findPart1(input, i)
		}
	}
	if resultPart1 != 514579 {
		t.Errorf("Expected 514579 but got %d", resultPart1)
	}
}
func TestExamplePart2(t *testing.T) {
	var input []int = []int{1721, 979, 366, 299, 675, 1456}
	var resultPart2 int
	var foundPart2 bool
	for i := range input {
		if !foundPart2 {
			resultPart2, foundPart2 = findPart2(input, i)
		}
	}
	if resultPart2 != 241861950 {
		t.Errorf("Expected 241861950 but got %d", resultPart2)
	}
}

func TestSolver(t *testing.T) {
	var input []int = []int{1721, 979, 366, 299, 675, 1456}
	part1, part2, err := Solver(input)
	if err != nil {
		t.Error("Didn't find a solution: ", err)
	}
	if part1 != 514579 && part2 != 241861950 {
		t.Errorf("Expected Part 1: 514579 but got %d\nExpected Part 2: 241861950 but got %d", part1, part2)
	}
}

func TestSolverErrors(t *testing.T) {
	var input []int = []int{1, 1, 2018}
	_, _, err := Solver(input)
	if err == nil {
		t.Error("Expected an error to return")
	}
	input = []int{1, 1, 1}
	_, _, err = Solver(input)
	if err == nil {
		t.Error("Expected an error to return")
	}

	input = []int{1, 2019, 1}
	_, _, err = Solver(input)
	if err == nil {
		t.Error("Expected an error to return")
	}
}

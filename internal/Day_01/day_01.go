package day01

import "errors"

func Solver(input []int) (int, int, error) {
	var resultPart1, resultPart2 int
	var foundPart1, foundPart2 bool
	var err error = nil
	for i := range input {
		if !foundPart1 {
			resultPart1, foundPart1 = findPart1(input, i)
		}
		if !foundPart2 {
			resultPart2, foundPart2 = findPart2(input, i)
		}
		if foundPart1 && foundPart2 {
			break
		}
	}
	if resultPart1+resultPart2 == 0 {
		err = errors.New("no solution was found")
	} else if resultPart1 == 0 {
		err = errors.New("no solution was found for part 1")
	} else if resultPart2 == 0 {
		err = errors.New("no solution was found for part 2")
	}
	return resultPart1, resultPart2, err

}

func findPart1(slice []int, currentIndex int) (int, bool) {
	for i := len(slice) - 1; i > currentIndex; i-- {
		if slice[i]+slice[currentIndex] == 2020 {
			return slice[i] * slice[currentIndex], true
		}
	}
	return 0, false
}

func findPart2(slice []int, currentIndex int) (int, bool) {
	for i := len(slice) - 1; i > currentIndex; i-- {
		for j := len(slice) - 1; j > currentIndex; j-- {
			if slice[i]+slice[currentIndex]+slice[j] == 2020 {
				return slice[i] * slice[currentIndex] * slice[j], true
			}
		}
	}
	return 0, false
}

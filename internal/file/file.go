package file

import (
	"bufio"
	"os"
)

func Read_Line(filepath string) (<-chan string, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	lines := make(chan string)

	go func() {
		defer close(lines)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	return lines, nil
}

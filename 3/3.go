package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, _ := readLines("3_input.txt")

	treesHit := 0
	//Assume vertical is always 1
	slope := 1
	currentX := 0

	for i, line := range lines {
		// rules := strings.Split(line[0], " ")
		if i%2 == 1 {
			continue
		}
		if line[currentX%len(line)] == '#' {
			treesHit += 1
		}
		currentX += slope
	}
	fmt.Println(treesHit)

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lines = append(lines, str)

	}
	return lines, scanner.Err()
}

//84 * 198 * 72 * 81 * 35

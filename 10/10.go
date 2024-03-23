package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type bagNumber struct {
	bag, number string
}

func main() {
	lines, _ := readLines("10_input.txt")
	sort.Ints(lines)
	fmt.Println(lines)
	oneDifference := 1
	threeDifference := 1
	for i, line := range lines[1:] {
		if line-lines[i] == 1 {
			oneDifference++
		} else if line-lines[i] == 3 {
			threeDifference++
		}
	}
	fmt.Println(oneDifference * threeDifference)

	//Star 1
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, str)

	}
	return lines, scanner.Err()
}

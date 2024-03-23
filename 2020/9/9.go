package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bagNumber struct {
	bag, number string
}

func main() {
	lines, _ := readLines("9_input.txt")

	//Star 1
	//Create set of 25 numbers
	// preambleBuffer := lines[:25]
	preambleBuffer := make(map[int]bool)
	for _, num := range lines[:25] {
		preambleBuffer[num] = true
	}
	//Iterate starting at number 26 - check against all combinations of set
	for i, num := range lines[25:] {
		valid := false
		for preambleNumber := range preambleBuffer {
			compliment := num - preambleNumber
			if compliment > 0 && compliment != preambleNumber && preambleBuffer[compliment] {
				valid = true
				break
			}
		}
		if valid {
			preambleBuffer[lines[i]] = false
			preambleBuffer[num] = true
		} else {
			fmt.Println(num)
			break
		}
	}

	//Star 2
	magicNumber := 1504371145

	//Loop through
	for i := range lines {
		j := i + 1
		for sumRange(lines[i:j]) < magicNumber {
			j += 1
		}
		if sumRange(lines[i:j]) == magicNumber {
			fmt.Println(findSmallestPlusLargest(lines[i:j]))
		}
	}
}

func findSmallestPlusLargest(array []int) int {
	smallest := array[0]
	largest := array[0]
	for _, val := range array {
		if val < smallest {
			smallest = val
		} else if val > largest {
			largest = val
		}
	}
	return smallest + largest
}

func sumRange(array []int) int {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	return sum
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

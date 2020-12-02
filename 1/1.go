package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, _ := readLinesToInt("1_input.txt")

	star2(lines)
}

func star1(lines []int) {
	m := make(map[int]bool)
	for _, num := range lines {
		m[num] = true
		cor := 2020 - num
		if m[cor] == true {
			fmt.Println(cor * num)
		}
	}
}

func star2(lines []int) {
	m := make(map[int]bool)
	for _, num := range lines {
		m[num] = true
	}
	for i, num := range lines {
		for _, num2 := range lines[i+1:] {
			cor := 2020 - (num + num2)
			if m[cor] == true {
				fmt.Println(cor * num * num2)
			}
		}
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.[
func readLinesToInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)

	}
	return lines, scanner.Err()
}

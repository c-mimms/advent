package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, _ := readLines("2_input.txt")

	//Count correct passwords
	count := 0
	for _, line := range lines {

		//Needs to be an easier way to do ths
		password := line[1]
		rules := strings.Split(line[0], " ")
		ruleRange := strings.Split(rules[0], "-")
		low, _ := strconv.Atoi(ruleRange[0])
		high, _ := strconv.Atoi(ruleRange[1])
		letter := rules[1]

		//Star 1
		// letterCount := strings.Count(password, letter)
		// if letterCount >= low && letterCount <= high {
		// 	count += 1
		// }

		//Star 2
		if (password[low] == letter[0]) != (password[high] == letter[0]) {
			count += 1
		}

	}
	fmt.Println(count)
	readLinesRegex("2_input.txt")
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lines = append(lines, strings.Split(str, ":"))

	}
	return lines, scanner.Err()
}

func readLinesRegex(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//1-3 g: nvgxtllpcgcgnmdqgg
	re1, err := regexp.Compile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		regexMatch := re1.FindStringSubmatch(str)
		lines = append(lines, regexMatch)
		fmt.Println(regexMatch[1:])

	}
	return lines, scanner.Err()
}

//Didn't end up using this
//4-5 c: ccssnccccc
// type passwordRecord struct {
// 	rule, password string
// }

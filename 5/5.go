package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, _ := readLines("5_input.txt")
	//FBFBBFFRLR
	//0101100  101

	highest := 0
	// line := "FBFBBFFRLR"
	var seatIds []int

	for _, line := range lines {
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")

		row, _ := strconv.ParseInt(line[0:7], 2, 0)
		column, _ := strconv.ParseInt(line[7:], 2, 0)

		checkSum := row*8 + column
		seatIds = append(seatIds, int(checkSum))
		if int(checkSum) > highest {
			highest = int(checkSum)
		}
	}
	sort.Ints(seatIds)

	for i, seat := range seatIds {
		// fmt.Println(seat + 1)
		if seat+2 == seatIds[i+1] {
			fmt.Println(seat + 1)

		}
	}
	fmt.Println(seatIds)

}

// byr:1994 hgt:152cm pid:198152466
// eyr:2022 ecl:hzl hcl:#4df239 iyr:2020

// ecl:grn
// eyr:2022
// byr:1968 iyr:2017 pid:044109096

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

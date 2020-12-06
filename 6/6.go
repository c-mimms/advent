package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, _ := readLinesGrouping("6_input.txt")
	fmt.Println(lines)

	summedQuestions := 0

	for _, group := range lines {

		questionMap := make([]int, 26)
		numPeeps := 0

		for _, person := range group {
			numPeeps += 1
			for _, letter := range person {
				questionMap[int(letter-'a')] += 1
			}
		}
		// fmt.Println(questionMap)

		uniqueQs := 0
		for _, q := range questionMap {
			if q == numPeeps {
				uniqueQs += 1
			}
		}
		// fmt.Println(uniqueQs)

		summedQuestions += uniqueQs
	}

	fmt.Println(summedQuestions)

}

// byr:1994 hgt:152cm pid:198152466
// eyr:2022 ecl:hzl hcl:#4df239 iyr:2020

// ecl:grn
// eyr:2022
// byr:1968 iyr:2017 pid:044109096

func readLinesGrouping(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)

	var id []string

	for scanner.Scan() {
		str := scanner.Text()
		if len(str) > 0 {
			id = append(id, str)
		} else {
			//End of group
			lines = append(lines, id)
			id = []string{}
		}
	}
	return lines, scanner.Err()
}

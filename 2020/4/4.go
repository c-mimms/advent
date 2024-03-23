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
	lines, _ := readLines("4_input.txt")

	//Count correct passwords
	// count := 0
	// for _, line := range lines {

	// 	//Needs to be an easier way to do ths
	// 	password := line[1]
	// 	rules := strings.Split(line[0], " ")
	// 	ruleRange := strings.Split(rules[0], "-")
	// 	low, _ := strconv.Atoi(ruleRange[0])
	// 	high, _ := strconv.Atoi(ruleRange[1])
	// 	letter := rules[1]
	// }
	count := 0
	// fmt.Println(lines)

	//Height regex
	hgtRx, _ := regexp.Compile(`([0-9]+)([a-z]+)`)
	hclRx, _ := regexp.Compile(`#([a-f0-9]+)`)
	eclRx, _ := regexp.Compile(`(amb|blu|brn|gry|grn|hzl|oth)`)

	for _, id := range lines {
		// if len(id) >= 7 {
		// fmt.Println(id)
		if byr, ok := id["byr"]; ok {
			if iyr, ok := id["iyr"]; ok {
				if eyr, ok := id["eyr"]; ok {
					if hgt, ok := id["hgt"]; ok {
						if hcl, ok := id["hcl"]; ok {
							if ecl, ok := id["ecl"]; ok {
								if pid, ok := id["pid"]; ok {

									byrI, _ := strconv.Atoi(byr)
									iyrI, _ := strconv.Atoi(iyr)
									eyrI, _ := strconv.Atoi(eyr)
									// fmt.Println(id)
									if byrI < 1920 || byrI > 2002 {
										continue
									}
									if iyrI < 2010 || iyrI > 2020 {
										continue
									}
									if eyrI < 2020 || eyrI > 2030 {
										continue
									}

									hclMatch := hclRx.FindStringSubmatch(hcl)
									if hclMatch == nil {
										continue
									} else {
										if len(hclMatch[1]) != 6 {
											continue
										}
									}

									eclMatch := eclRx.FindStringSubmatch(ecl)
									if eclMatch == nil {
										continue
									}

									hgtMatch := hgtRx.FindStringSubmatch(hgt)
									if hgtMatch == nil {
										continue
									}
									// fmt.Println(hgtMatch)

									hgtIi, _ := strconv.Atoi(hgtMatch[1])
									if hgtMatch[2] == "in" {
										if !(hgtIi >= 59 && hgtIi <= 76) {
											continue
										}
									} else if hgtMatch[2] == "cm" {
										if !(hgtIi >= 150 && hgtIi <= 193) {
											continue
										}
									}
									if len(pid) != 9 {
										continue
									}

									count += 1
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

// byr:1994 hgt:152cm pid:198152466
// eyr:2022 ecl:hzl hcl:#4df239 iyr:2020

// ecl:grn
// eyr:2022
// byr:1968 iyr:2017 pid:044109096

func readLines(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []map[string]string
	scanner := bufio.NewScanner(file)
	// var id []string
	id := make(map[string]string)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) > 0 {
			for _, field := range strings.Split(str, " ") {
				keyVal := strings.Split(field, ":")
				id[keyVal[0]] = keyVal[1]
			}
		} else {
			//End of passport
			lines = append(lines, id)
			id = make(map[string]string)
		}
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

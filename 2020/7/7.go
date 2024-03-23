package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bagNumber struct {
	bag, number string
}

func main() {
	lines, _ := readLines("7_input.txt")

	//This makes a regular map, I am going to make an inverted map instead
	//Ended up using this for part 2 but added a capture group for the number

	rulesMap := make(map[string][]bagNumber)

	for _, line := range lines {
		//dark maroon bags contain 5 dotted orange bags, 5 faded beige bags, 3 wavy salmon bags, 5 dim lavender bags.
		splitLine := strings.Split(line, " bags contain ")

		rulesRegex, _ := regexp.Compile(`([0-9]) (.+) bags?.?`)

		contains := strings.Split(splitLine[1], ", ")
		var rules []bagNumber

		for _, bag := range contains {
			rulesMatch := rulesRegex.FindStringSubmatch(bag)
			if rulesMatch != nil {
				bagNum := bagNumber{rulesMatch[2], rulesMatch[1]}
				rules = append(rules, bagNum)
			}
		}

		rulesMap[splitLine[0]] = rules
	}

	//This makes an inverted map, used for part 1
	// rulesMap := make(map[string][]string)

	// for _, line := range lines {
	// 	//dark maroon bags contain 5 dotted orange bags, 5 faded beige bags, 3 wavy salmon bags, 5 dim lavender bags.
	// 	splitLine := strings.Split(line, " bags contain ")

	// 	rulesRegex, _ := regexp.Compile(`[0-9] (.+) bags?.?`)

	// 	contains := strings.Split(splitLine[1], ", ")
	// 	// var rules []string

	// 	for _, bag := range contains {
	// 		rulesMatch := rulesRegex.FindStringSubmatch(bag)
	// 		if rulesMatch != nil {
	// 			bagList, _ := rulesMap[rulesMatch[1]]
	// 			bagList = append(bagList, splitLine[0])
	// 			rulesMap[rulesMatch[1]] = bagList
	// 		}
	// 	}
	// }

	// //Recursively count paths through map - Star 1
	// counted := make(map[string]bool)
	// fmt.Println(countPathsExcludeCounted("shiny gold", counted, rulesMap))

	//Recursively count paths through map - Star 2
	fmt.Println(countPaths("shiny gold", rulesMap) - 1)
}

func countPaths(name string, routeMap map[string][]bagNumber) int {
	routes := routeMap[name]
	count := 1
	for _, route := range routes {
		num, _ := strconv.Atoi(route.number)
		fmt.Println(name, "requires", num, route.bag)
		count += num * countPaths(route.bag, routeMap)
	}
	fmt.Println(name, "contains", count)
	return count
}

func countPathsExcludeCounted(name string, counted map[string]bool, routeMap map[string][]string) int {
	count := 0
	routes := routeMap[name]
	for _, route := range routes {
		if !counted[route] {
			counted[route] = true
			count += 1 + countPathsExcludeCounted(route, counted, routeMap)
		}
	}
	return count
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

package main

import (
	"strconv"
	"strings"
)

func solve1() {
	inputLocations := readFileToArray("input6.txt")
	locations := parseInput(inputLocations)

	var mapp [400][400]int
	locationCount := make([]int, len(locations))

	//Find the closest location for each coordinate
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			closest := findClosestLocation(i, j, locations)
			mapp[i][j] = closest
			locationCount[closest] += 1
		}
	}

	//Any edge coordinate location is infinite
	for i := 0; i < 400; i++ {
		leftEdge := mapp[i][0]
		rightEdge := mapp[i][399]
		topEdge := mapp[0][i]
		bottomEdge := mapp[399][i]
		locationCount[leftEdge] = 0
		locationCount[rightEdge] = 0
		locationCount[topEdge] = 0
		locationCount[bottomEdge] = 0
	}

	highestCount := 0
	highestLocation := -1
	for i, count := range locationCount {
		if count > highestCount {
			highestCount = count
			highestLocation = i
		}
	}
	printVal(highestCount)
	printVal(highestLocation)
}

func solve2() {
	inputLocations := readFileToArray("input6.txt")
	locations := parseInput(inputLocations)

	// var mapp [400][400]int
	// locationCount := make([]int, len(locations))
	numOfCloseRegions := 0

	//Find the closest location for each coordinate
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			sumDistance := 0
			for _, location := range locations {
				sumDistance += taxiDistance(i, j, location[0], location[1])
			}
			if sumDistance < 10000 {
				numOfCloseRegions += 1
			}
		}
	}
	printVal(numOfCloseRegions)
}

func parseInput(input []string) [][]int {
	var locations [][]int
	for _, location := range input {
		splitLocation := splitAndTrim(location, ",", " ")
		x, _ := strconv.Atoi(splitLocation[0])
		y, _ := strconv.Atoi(splitLocation[1])
		locations = append(locations, []int{x, y})
	}
	return locations
}

func findClosestLocation(x int, y int, locations [][]int) int {
	closest := -1
	closestDist := 400
	for i, location := range locations {
		distance := taxiDistance(x, y, location[0], location[1])
		if distance < closestDist {
			closestDist = distance
			closest = i
		}
	}
	return closest
}

func taxiDistance(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func splitAndTrim(input string, sep string, remove string) []string {
	splits := strings.Split(input, sep)
	var trimmed []string
	for _, line := range splits {
		line = strings.Trim(line, remove)
		trimmed = append(trimmed, line)
	}
	return trimmed
}

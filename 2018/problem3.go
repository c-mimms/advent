package main

import (
	"strconv"
	"strings"
)

func solve1() {
	elfClaims := readFileToArray("input3.txt")

	cloth := make([][]uint8, 1200)
	for i := range cloth {
		cloth[i] = make([]uint8, 1200)
	}

	for _, claim := range elfClaims {
		parts := strings.Split(claim, " ")
		start := strings.Split(parts[2][:len(parts[2])-1], ",")
		size := strings.Split(parts[3], "x")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])

		sizeX, _ := strconv.Atoi(size[0])
		sizeY, _ := strconv.Atoi(size[1])

		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				cloth[startX+i][startY+j] += 1
			}
		}
	}

	count := 0
	for _, row := range cloth {
		for _, val := range row {
			if val > 1 {
				count += 1
			}
		}
	}
	printVal(count)
}

func solve2() {
	elfClaims := readFileToArray("input3.txt")

	cloth := make([][]uint8, 1200)
	for i := range cloth {
		cloth[i] = make([]uint8, 1200)
	}

	for _, claim := range elfClaims {
		parts := strings.Split(claim, " ")
		start := strings.Split(parts[2][:len(parts[2])-1], ",")
		size := strings.Split(parts[3], "x")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])

		sizeX, _ := strconv.Atoi(size[0])
		sizeY, _ := strconv.Atoi(size[1])

		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				cloth[startX+i][startY+j] += 1
			}
		}
	}

	for num, claim := range elfClaims {
		parts := strings.Split(claim, " ")
		start := strings.Split(parts[2][:len(parts[2])-1], ",")
		size := strings.Split(parts[3], "x")

		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])

		sizeX, _ := strconv.Atoi(size[0])
		sizeY, _ := strconv.Atoi(size[1])

		freeClaim := true
		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				if cloth[startX+i][startY+j] > 1 {
					freeClaim = false
				}
			}
		}
		if freeClaim {
			printVal(num + 1)
			return
		}
	}
}

package main

import (
	"fmt"
)

func solve1() {
	boxLables := readFileToArray("input2.txt")

	twoCount := 0
	threeCount := 0

	for _, label := range boxLables {
		isTwo := false
		isThree := false
		letterCounts := make(map[rune]int)
		for _, letter := range label {
			letterCounts[letter] += 1
		}
		for _, val := range letterCounts {
			if val == 2 {
				isTwo = true
			} else if val == 3 {
				isThree = true
			}
		}
		if isTwo {
			twoCount += 1
		}
		if isThree {
			threeCount += 1
		}
	}
	fmt.Println(twoCount * threeCount)
}

func solve2() {
	boxLables := readFileToArray("input2.txt")

	for i, label := range boxLables {
		for _, label2 := range boxLables[i:] {
			var answer []byte
			for j, _ := range label {
				if label2[j] == label[j] {
					answer = append(answer, label2[j])
				}
			}
			if len(answer) == len(label2)-1 {
				fmt.Println(string(answer))
			}
		}
	}
}

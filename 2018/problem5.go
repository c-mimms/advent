package main

import "strings"

func solve1() {
	inputPolymer := []byte(readFileToArray("input5.txt")[0])
	polymer := removeDupes(inputPolymer)
	printVal(len(polymer))
}

func solve2() {
	inputPolymer := readFileToArray("input5.txt")[0]
	minLength := 10000
	for i := 'A'; i <= 'Z'; i++ {
		polymerRemoved := strings.Replace(inputPolymer, string(i), "", -1)
		polymerRemoved = strings.Replace(polymerRemoved, string(i+32), "", -1)
		polymer := removeDupes([]byte(polymerRemoved))
		if len(polymer) < minLength {
			minLength = len(polymer)
		}
	}
	printVal(minLength)
}

func removeDupes(s []byte) []byte {
	length := len(s)
	for i, char := range s {
		if i == length-1 {
			continue
		}
		if areCharsSameButDifferent(char, s[i+1]) {
			var newString []byte
			newString = append(newString, s[:i]...)
			newString = append(newString, s[i+2:]...)
			return removeDupes(newString)
		}
	}
	return s
}

func areCharsSameButDifferent(a byte, b byte) bool {
	if a+32 == b {
		return true
	}
	if a-32 == b {
		return true
	}
	return false
}

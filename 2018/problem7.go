package main

import (
	"sort"
	"strings"
)

const BaseStepTime = 60
const NumWorkers = 5

func solve2() {
	finalOrder, instructions := solve1()
	finalOrder = append(finalOrder, "V")
	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for _, char := range alphabet {
		printVal(getTimeForStep(char))
	}
	printVal(finalOrder)
	printVal(instructions)
}

func getTimeForStep(step byte) int {
	return int(step - 'A' + 61)
}

func solve1() ([]string, map[string][]string) {
	inputSentences := readFileToArray("input7.txt")
	instructions := make(map[string][]string)
	depList := make(map[string]bool)
	for _, sentence := range inputSentences {
		words := strings.Split(sentence, " ")
		required := words[1]
		step := words[7]
		depList[required] = true
		precedents := instructions[step]
		precedents = append(precedents, required)
		instructions[step] = precedents
	}
	// printVal(instructions)
	// firsts := getFirsts(instructions, depList)
	// printVal(firsts)
	// printVal(depList)

	var allSteps []string
	for step, _ := range depList {
		allSteps = append(allSteps, step)
	}
	sort.Strings(allSteps)
	// printVal(allSteps)

	var finalOrder []string
	for len(finalOrder) < 25 {
		for i, step := range allSteps {
			requirements, ok := instructions[step]
			requirementsFulfilled := checkReq(finalOrder, requirements)
			if !ok || requirementsFulfilled {
				finalOrder = append(finalOrder, step)
				allSteps = remove(allSteps, i)
				break
			}
		}
	}

	printVal(strings.Join(finalOrder, ""))
	return finalOrder, instructions
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func checkReq(finalOrder []string, requirements []string) bool {
	for _, requiredStep := range requirements {
		if !contains(finalOrder, requiredStep) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(a, e) == 0 {
			return true
		}
	}
	return false
}

func getFirsts(instructions map[string][]string, depList map[string]bool) map[string]string {
	firsts := make(map[string]string)
	for letter, _ := range depList {
		_, ok := instructions[letter]
		if !ok {
			firsts[letter] = letter
		}
	}
	return firsts
}

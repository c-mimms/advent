package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagNumber struct {
	bag, number string
}

func main() {
	lines, _ := readLines("8_input.txt")

	//Star 1

	//Star 2
	for changeInstruction, _ := range lines {
		accumulator := 0
		programCounter := 0
		step := 0
		visitedInstructions := make(map[int]bool)
		for !visitedInstructions[programCounter] && programCounter != len(lines)-1 {
			visitedInstructions[programCounter] = true
			operation, argument := parseInstruction(lines[programCounter])
			if step == changeInstruction {
				if operation == "jmp" {
					operation = "nop"
				} else if operation == "nop" {
					operation = "jmp"
				}
			}
			switch operation {
			case "jmp":
				programCounter += argument
			case "acc":
				accumulator += argument
				programCounter += 1
			case "nop":
				programCounter += 1
			}
			step += 1
		}
		if programCounter == len(lines)-1 {
			fmt.Println(accumulator)
			break
		}
	}

}

func parseInstruction(instruction string) (string, int) {
	splitInstructions := strings.Split(instruction, " ")
	arg, _ := strconv.Atoi(splitInstructions[1])
	return splitInstructions[0], arg
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

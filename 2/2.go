package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, _ := readLines("2_input.txt")

	//Count correct passwords
	for _, line := range lines {
		// password := line[1]
		rules := strings.Split(line[0], " ")
		fmt.Println(rules[0])
	}

}

//4-5 c: ccssnccccc
type passwordRecord struct {
	rule, password string
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		// num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, strings.Split(str, ":"))

	}
	return lines, scanner.Err()
}

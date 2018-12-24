package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solve1()
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileToArray(filename string) []string {
	f, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(f)
	fileStrings := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fileStrings = append(fileStrings, line)
	}
	f.Close()
	return fileStrings
}

func printVal(a ...interface{}) {
	fmt.Printf("%v\n", a)
}

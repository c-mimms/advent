// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func solve1() {
	f, err := os.Open("input1.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	frequency := 0

	for scanner.Scan() {
		change := scanner.Text()
		operator := change[0]
		number, err := strconv.Atoi(change[1:])
		check(err)
		fmt.Printf("%c , %v\n", operator, number)
		if operator == '+' {
			frequency += number
		} else if operator == '-' {
			frequency -= number
		}
	}
	fmt.Printf("%v\n", frequency)

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).
	f.Close()
}

func solve2() {
	dat, err := ioutil.ReadFile("input1.txt")
	check(err)
	b := bytes.NewReader(dat)
	scanner := bufio.NewScanner(b)

	visited := make(map[int]bool)

	frequency := 0

	for visited[frequency] != true {
		visited[frequency] = true
		if !scanner.Scan() {
			//All three of these work, but not alone
			b.Seek(0, 0)
			// b.Reset(dat)
			// b = bytes.NewReader(dat)
			scanner = bufio.NewScanner(b)
			scanner.Scan()
		}
		change := scanner.Text()
		operator := change[0]
		number, err := strconv.Atoi(change[1:])
		check(err)
		if operator == '+' {
			frequency += number
		} else if operator == '-' {
			frequency -= number
		}
	}
	fmt.Printf("%v\n", frequency)
}

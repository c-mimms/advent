package main

import (
	"bufio"
	"fmt"
	"os"
)

type bagNumber struct {
	bag, number string
}

func main() {

	//Star 1
	lines, _ := readLines("11_input.txt")
	// prettyPrint(lines)

	var lines2 []string
	changes := true

	for changes {
		lines2 = make([]string, 0)
		changes = false

		for i, line := range lines {
			var cloneline string
			for j, char := range line {
				//Count adjacentsies
				count := countAdjacent(i, j, lines)
				// fmt.Println(count)
				if char == 'L' && count == 0 {
					cloneline = cloneline + string('#')
					changes = true
				} else if char == '#' && count > 4 { //Added 1 to count to account for the character itself
					cloneline = cloneline + string('L')
					changes = true
				} else {
					cloneline = cloneline + string(char)
				}
			}
			lines2 = append(lines2, cloneline)
		}
		prettyPrint(lines2)
		fmt.Println(" ")
		lines = lines2
	}

	fmt.Println(countOccupied(lines))

	// oneDifference := 1
	// threeDifference := 1
	// for i, line := range lines[1:] {
	// 	if line-lines[i] == 1 {
	// 		oneDifference++
	// 	} else if line-lines[i] == 3 {
	// 		threeDifference++
	// 	}
	// }
	// fmt.Println(oneDifference * threeDifference)

	//Star 2

	//

}

func countAdjacent(i int, j int, lines []string) int {
	count := 0
	for _, line2 := range lines[max(i-1, 0):min(i+2, len(lines))] {
		// fmt.Println(line2[max(j-1, 0):min(j+2, len(line2))])
		for _, char2 := range line2[max(j-1, 0):min(j+2, len(line2))] {
			if char2 == '#' {
				count++
			}
		}
	}
	return count
}

func countOccupied(slice []string) int {
	count := 0
	for _, line := range slice {
		for _, char := range line {
			if char == '#' {
				count++
			}
		}
	}
	return count
}

func prettyPrint(slice []string) {
	for _, line := range slice {
		fmt.Println(line)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		// str, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, scanner.Text())

	}
	return lines, scanner.Err()
}

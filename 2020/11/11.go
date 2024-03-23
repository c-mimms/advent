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
		// for z := 0; z < 2; z++ {
		lines2 = make([]string, 0)
		changes = false

		for i, line := range lines {
			var cloneline string
			for j, char := range line {
				//Count adjacentsies (star 1)
				// count := countAdjacent(i, j, lines)

				//Count visible (Star 2)
				count := countVisible(j, i, lines)
				// fmt.Println(j, ",", i, " has ", count)

				if char == 'L' && count == 0 {
					cloneline = cloneline + string('#')
					changes = true
				} else if char == '#' && count > 4 { //Added 1 to count to account for the character itself (Add 1 for star 2)
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
	fmt.Println(changes)
}

func countVisible(startX int, startY int, lines []string) int {
	visibleOccupiedSeats := 0
	for _, xChange := range []int{-1, 0, 1} {
		for _, yChange := range []int{-1, 0, 1} {
			if xChange == 0 && yChange == 0 {
				continue
			}
			x := startX + xChange
			y := startY + yChange
			for inBounds(x, y, lines) {
				if lines[y][x] == 'L' {
					// fmt.Println("saw empty seat in direction ", xChange, yChange, "from", startX, startY)
					break
				} else if lines[y][x] == '#' {
					// fmt.Println("saw full seat in direction ", xChange, yChange, "from", startX, startY)
					visibleOccupiedSeats++
					break
				}
				x += xChange
				y += yChange
			}
		}
	}
	return visibleOccupiedSeats
}

//Returns true if x and y are within the array bounds
func inBounds(x int, y int, lines []string) bool {
	if y >= 0 && y < len(lines) && x >= 0 && x < len(lines[0]) {
		return true
	}
	return false
}

func countAdjacent(i int, j int, lines []string) int {
	count := 0
	for _, line2 := range lines[max(i-1, 0):min(i+2, len(lines))] {
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

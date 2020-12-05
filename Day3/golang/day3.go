package main

import (
	"bufio"
	"fmt"
	"os"
)

const TreeRune = "#"

func main() {
	//tobogganRun,err := readLines("../data/test.map")
	tobogganRun, err := readLines("../data/input.map")
	check(err)

	startX, startY := 0, 0
	moveRight, moveDown := 3, 1

	numberOfTrees := SolveTobogganRun(tobogganRun, startX, startY, moveRight, moveDown)
	fmt.Println("Success, you hit ", numberOfTrees, " trees on your way down.")

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	productOfTrees := 1
	for _, slope := range slopes {
		numberOfTrees := SolveTobogganRun(tobogganRun, startX, startY, slope[0], slope[1])
		productOfTrees *= numberOfTrees
		fmt.Printf("With slope %v, you hit %d trees on your way down.\n", slope, numberOfTrees)
	}

	fmt.Printf("Total product of all trees hit: %d\n", productOfTrees)
}

func SolveTobogganRun(field []string, startX int, startY int, moveRight int, moveDown int) int {
	treeCount := 0

	row := startX
	col := startY

	for row < len(field) {
		if hasTreeAt(field, row, col) {
			treeCount++
		}

		col = (col + moveRight) % len(field[row])
		row = row + moveDown

	}

	return treeCount
}

func hasTreeAt(field []string, row int, col int) bool {
	value := field[row][col]

	if value == '#' {
		return true
	}

	return false
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
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

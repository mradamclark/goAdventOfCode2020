package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type seatInfo struct {
	row, seat int
}

func main() {
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day5/data/simpleTest.dat"
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day5/data/test.dat"
	fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day5/data/input.dat"

	ticketData, err := readLines(fileWithPath)
	check(err)

	currentSeatId := 0
	for _, seatReference := range ticketData {
		seatLoc := findSeatLocation(strings.ToUpper(seatReference))
		seatID := (seatLoc.row * 8) + seatLoc.seat
		fmt.Printf("SeatReference: %s %v - %d\n", seatReference, seatLoc, seatID)
		if seatID > currentSeatId {
			currentSeatId = seatID
		}
	}

	fmt.Println("Maximum seat ID seen is: ", currentSeatId)
}

func findSeatLocation(ticketReference string) seatInfo {
	rowReference := ticketReference[0:7]
	row := bsp_findpoint(rowReference, 0, 127, 'F')

	seatReference := ticketReference[7:]
	seat := bsp_findpoint(seatReference, 0, 7, 'L')

	return seatInfo{row: row, seat: seat}
}

func bsp_findpoint(seatReference string, min int, max int, lower byte) int {
	//fmt.Printf("min:%3d, max:%3d - %s\n", min, max, seatReference)

	if len(seatReference) == 1 {
		if seatReference == string(lower) {
			return min
		} else {
			return max
		}
	}

	choice := seatReference[0]
	var newMin, newMax int
	if choice == lower {
		newMax = ((max - min) / 2) + min
		newMin = min
	} else {
		newMax = max
		newMin = max - ((max - min) / 2)
	}

	return bsp_findpoint(seatReference[1:], newMin, newMax, lower)
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

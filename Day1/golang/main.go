package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := readLines("input.txt")
	check(err)

	//test input.
	//data = []string{"1721", "979", "366", "299", "675", "1456"}

	fmt.Println("Part 1...")
	solvePart1(data)

	fmt.Println("Part 2...")
	solvePart2(data)

}

func solvePart1(data []string) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			first, _ := strconv.Atoi(data[i])
			second, _ := strconv.Atoi(data[j])

			var addResult = first + second

			if 2020 == addResult {
				fmt.Printf("Success %d * %d = %d\n", first, second, first*second)
				return
			}
		}
	}

	fmt.Println("Failure, nothing in this data set")
}

func solvePart2(data []string) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				first, _ := strconv.Atoi(data[i])
				second, _ := strconv.Atoi(data[j])
				third, _ := strconv.Atoi(data[k])

				var addResult = first + second + third

				if 2020 == addResult {
					fmt.Printf("Success %d * %d * %d = %d\n\n", first, second, third, first*second)
					return
				}
			}
		}
	}

	fmt.Println("Failure, nothing in this data set")
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

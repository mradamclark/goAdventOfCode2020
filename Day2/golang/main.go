package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type policy struct {
	min   int
	max   int
	value string
}

func newPolicy(policyDefinition string) *policy {
	var p policy
	_, err := fmt.Sscanf(policyDefinition, "%d-%d %s", &p.min, &p.max, &p.value)
	check(err)

	return &p
}

func main() {
	pwdInput, err := readLines("../data/input.txt")
	check(err)

	fmt.Println("\n\nPart One - Policy Violations: ", countValidPasswords(pwdInput))
}

func countValidPasswords(input []string) int {
	passes := 0

	for _, entry := range input {
		rule, pwd, err := parseInputEntry(entry)
		check(err)

		if passwordValidator(pwd, *rule) {
			passes += 1
		}

		//fmt.Printf("\nPolicy: %v, Pwd: %s", rule, pwd)
	}
	return passes
}

func passwordValidator(pwd string, rule policy) bool {

	count := strings.Count(pwd, rule.value)

	if count < rule.min || count > rule.max {
		return false
	}

	return true
}

func parseInputEntry(entry string) (*policy, string, error) {
	parts := strings.Split(entry, ":")

	if len(parts) == 2 {
		return newPolicy(parts[0]), parts[1], nil
	}

	return nil, "", errors.New("invalid entry")
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

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
	_, err := fmt.Sscanf(policyDefinition, "%d-%d %v", &p.min, &p.max, &p.value)
	check(err)

	return &p
}

type validator func(pwd string, rule policy) bool

func main() {
	//pwdInput, err := readLines("../data/test.txt")
	pwdInput, err := readLines("../data/input.txt")
	check(err)

	fmt.Println(
		"\n\nPart One - Valid Passwords (Schema 1): ",
		countValidPasswords(pwdInput, passwordValidatorPart1))

	fmt.Println(
		"\nPart Two - Valid Passwords (Schema 2)",
		countValidPasswords(pwdInput, passwordValidatorPart2))
}

func countValidPasswords(input []string, tester validator) int {
	passes := 0

	for _, entry := range input {
		rule, pwd, err := parseInputEntry(entry)
		check(err)

		if tester(pwd, *rule) {
			passes += 1
		}
	}
	return passes
}

func passwordValidatorPart1(pwd string, rule policy) bool {
	count := strings.Count(pwd, rule.value)

	if count < rule.min || count > rule.max {
		return false
	}

	return true
}

func passwordValidatorPart2(pwd string, rule policy) bool {
	first := false
	second := false

	if len(pwd) >= rule.min {
		//fmt.Print("\nFirst chara: ",string(pwd[rule.min]))
		first = rule.value == string(pwd[rule.min])
	}

	if len(pwd) >= rule.max {
		//fmt.Print(", Second char: ",string(pwd[rule.max]))
		second = rule.value == string(pwd[rule.max])
	}

	//fmt.Printf("\nPolicy: %v, Pwd: %s, %v", rule, pwd, (first && !second) || (!first && second))
	return (first && !second) || (!first && second)
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

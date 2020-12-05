package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var requiredFields = []string{
	"pid",
	//"cid", //countryID is not required
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl"}

func main() {
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/test.dat"
	fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/input.dat"
	file, err := os.Open(fileWithPath)
	if err != nil {
		log.Fatal(err)
	}

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	passports := parseRawData(rawData)
	validPassports := 0

	fmt.Println("Number of Passports: ", len(passports))

	valid := false
	for _, passport := range passports {
		valid = validatePassport(passport)
		if valid {
			validPassports++
		}
	}

	fmt.Println("Number of valid passports: ", validPassports)
}

func validatePassport(passport map[string]string) bool {
	keys := getMapKeys(passport)
	return containsAllKeys(keys, requiredFields)
}

func containsAllKeys(keys []string, requiredFields []string) bool {
	hasAllKeys := true

	for _, rf := range requiredFields {
		keyFound := false
		for _, k := range keys {
			if k == rf {
				keyFound = true
			}
		}
		hasAllKeys = hasAllKeys && keyFound
	}

	return hasAllKeys
}

func getMapKeys(passport map[string]string) []string {
	keys := make([]string, len(passport))
	i := 0
	for k := range passport {
		keys[i] = k
		i++
	}
	return keys
}

func parseRawData(data []byte) []map[string]string {
	var passports []map[string]string

	records := bytes.SplitN(data, []byte{'\n', '\n'}, -1)

	for _, rec := range records {
		fields := bytes.Split(bytes.ReplaceAll(rec, []byte{'\n'}, []byte{' '}), []byte{' '})
		passport := make(map[string]string)
		for _, field := range fields {
			keyValuePair := strings.Split(string(field), ":")
			passport[keyValuePair[0]] = keyValuePair[1]
		}

		//usage of append causes lots of memory allocations to occur
		//should find a more efficient way to do this
		passports = append(passports, passport)
	}

	return passports
}

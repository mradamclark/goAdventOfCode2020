package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/test.dat"
	fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/input.dat"
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/invalid.dat"
	//fileWithPath := "/Users/adamclark/dev/projects/goAdventOfCode/Day4/data/valid.dat"
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

	for _, passport := range passports {
		//"cid", //countryID is not required
		valid := validatePassport(passport, []string{"pid", "byr", "iyr", "eyr", "hgt", "hcl", "ecl"})

		if valid {
			validPassports++
		}
	}
	fmt.Println("Part 1 Validation Result: ", validPassports)

	validPassports = 0
	for _, passport := range passports {
		valid, _ := validatePassportV2(passport)
		if valid {
			validPassports++
		}
	}
	fmt.Println("Part 2 Validation Result: ", validPassports)

}

func validatePassportV2(passport map[string]string) (bool, map[string]bool) {
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	//cid (Country ID) - ignored, missing or not.

	if len(passport) < 7 {
		return false, nil
	}

	results := make(map[string]bool)
	var valid = true

	if val, ok := passport["byr"]; ok {
		results["byr"] = validateYear(val, 1920, 2002)
		valid = valid && results["byr"]
	} else {
		valid = false
	}

	if val, ok := passport["iyr"]; valid && ok {
		results["iyr"] = validateYear(val, 2010, 2020)
		valid = valid && results["iyr"]
	} else {
		valid = false
	}

	if val, ok := passport["eyr"]; valid && ok {
		results["eyr"] = validateYear(val, 2020, 2030)
		valid = valid && results["eyr"]
	} else {
		valid = false
	}

	if val, ok := passport["hgt"]; valid && ok {
		results["hgt"] = validateHeight(val, 150, 193, 59, 76)
		valid = valid && results["hgt"]
	} else {
		valid = false
	}

	if val, ok := passport["hcl"]; valid && ok {
		results["hcl"] = validateHairColour(val)
		valid = valid && results["hcl"]
	} else {
		valid = false
	}

	if val, ok := passport["ecl"]; valid && ok {
		results["ecl"] = validateEyeColour(val, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
		valid = valid && results["ecl"]
	} else {
		valid = false
	}

	if val, ok := passport["pid"]; valid && ok {
		results["pid"] = validatePassportID(val)
		valid = valid && results["pid"]
	} else {
		valid = false
	}

	return valid, results
}

func validatePassportID(value string) bool {
	if len(value) != 9 {
		return false
	}

	_, err := strconv.Atoi(value)
	return err == nil
}

func validateHairColour(value string) bool {
	if !strings.HasPrefix(value, "#") {
		return false
	}

	intermediate := strings.TrimPrefix(value, "#")
	if len(intermediate) != 6 {
		return false
	}

	_, err := hex.DecodeString(intermediate)
	return err == nil
}

func validateEyeColour(value string, validEyeColours []string) bool {
	for _, validColour := range validEyeColours {
		if value == validColour {
			return true
		}
	}

	return false
}

func validateHeight(value string, cmMin int, cmMax int, inMin int, inMax int) bool {
	if strings.HasSuffix(value, "cm") {
		number, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
		if err == nil && number >= cmMin && number <= cmMax {
			return true
		}
	} else if strings.HasSuffix(value, "in") {
		number, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
		if err == nil && number >= inMin && number <= inMax {
			return true
		}
	}

	return false
}

func validateYear(value string, min int, max int) bool {
	intermediate, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return intermediate >= min && intermediate <= max
}

func validatePassport(passport map[string]string, requiredFields []string) bool {
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
			if keyValuePair[0] != "cid" {
				passport[keyValuePair[0]] = keyValuePair[1]
			}
		}

		//usage of append causes lots of memory allocations to occur
		//should find a more efficient way to do this
		passports = append(passports, passport)
	}

	return passports
}

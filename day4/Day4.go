package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PotentialPassport struct {
	fields map[string]string
}

func createPotentialPassport(lineOfInput string) PotentialPassport {
	words := strings.Fields(lineOfInput)
	fields := make(map[string]string)
	for _, passportField := range words {
		fieldName := strings.Split(passportField, ":")[0]
		fieldValue := strings.Split(passportField, ":")[1]
		fields[fieldName] = fieldValue
	}
	pp := PotentialPassport{
		fields: fields,
	}

	return pp
}

func (pp *PotentialPassport) isValid() bool {
	var fieldNames = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" /*,"cid"*/}
	for _, fieldName := range fieldNames {
		fieldValue, present := pp.fields[fieldName]
		if !present || !isFieldValid(fieldName, fieldValue) {
			fmt.Print("Failing password: ", pp.fields)
			fmt.Printf("Due to field: %s : %s\n", fieldName, fieldValue)
			return false
		}
	}
	return true
}

func isFieldValid(fieldName string, fieldValue string) bool {
	var isValid bool
	switch fieldName {
	case "byr":
		birthYear, _ := strconv.Atoi(fieldValue)
		isValid = 1920 <= birthYear && birthYear <= 2002
	case "iyr":
		issueYear, _ := strconv.Atoi(fieldValue)
		isValid = 2010 <= issueYear && issueYear <= 2020
	case "eyr":
		expireYear, _ := strconv.Atoi(fieldValue)
		isValid = 2020 <= expireYear && expireYear <= 2030
	case "hgt":
		heightType := fieldValue[len(fieldValue)-2:]
		heightMagnitude := fieldValue[0 : len(fieldValue)-2]
		isValid = isHeightValid(heightType, heightMagnitude)
	case "hcl":
		var validHairColor = regexp.MustCompile(`^#[\d\w]{6}$`)
		isValid = validHairColor.MatchString(fieldValue)
	case "ecl":
		var validEyeColor = regexp.MustCompile(`^((amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth))$`)
		isValid = validEyeColor.MatchString(fieldValue)
	case "pid":
		var validPassportID = regexp.MustCompile(`^\d{9}$`)
		isValid = validPassportID.MatchString(fieldValue)
	default:
		isValid = true
	}

	return isValid
}

func isHeightValid(units string, magnitudeString string) bool {
	switch units {
	case "cm":
		magnitude, _ := strconv.Atoi(magnitudeString)
		return 150 <= magnitude && magnitude <= 193
	case "in":
		magnitude, _ := strconv.Atoi(magnitudeString)
		return 59 <= magnitude && magnitude <= 76
	default:
		return false
	}
}

func Solve(potentialPassports []string) int {
	count := 0
	for _, potentialPassport := range potentialPassports {
		passport := createPotentialPassport(potentialPassport)
		if passport.isValid() {
			count++
		}
	}
	return count
}

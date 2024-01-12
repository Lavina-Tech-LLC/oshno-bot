package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ValidateDateBirth(dateBirth string) bool {
	splitDate := strings.Split(dateBirth, "-")
	if len(splitDate) != 3 {
		return false
	}
	day, err := strconv.Atoi(splitDate[0])
	if err != nil || day > 31 {
		fmt.Println(err, "error in check day")
		return false
	}

	month, err := strconv.Atoi(splitDate[1])
	if err != nil || month > 12 {
		fmt.Println(err, "error in check month")
		return false
	}

	year, err := strconv.Atoi(splitDate[2])
	if err != nil {
		fmt.Println(err, "error in check month")
		return false
	}
	if year < 1900 || year > 2005 {
		return false
	}
	return true
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidateString(input string) bool {
	for _, char := range input {
		if rune(char) == ' ' {
			continue
		}
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

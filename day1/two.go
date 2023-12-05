package day1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Two(inputFile string) int {

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var total int = 0
	for _, line := range strings.Split(string(data), "\n") {
		nr, err := strconv.Atoi(getFirstNumWithWord(line) + getLastNumWithWord(line))
		if err != nil {
			log.Fatal(err)
		}
		total += nr
	}

	return total
}

var numberWords = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func getLastNumWithWord(line string) string {

	revLine := reverseString(line)

	pattern, err := regexp.Compile("([0-9]|(?:eno|owt|eerht|ruof|evif|xis|neves|thgie|enin))")
	if err != nil {
		log.Fatal(err)
	}
	maybeNumber := pattern.FindString(revLine)

	if len(maybeNumber) > 1 {
		return numberWords[reverseString(maybeNumber)]
	}
	return maybeNumber
}

func reverseString(line string) string {
	chars := []rune(line)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func getFirstNumWithWord(line string) string {
	pattern, err := regexp.Compile("([0-9]|(?:one|two|three|four|five|six|seven|eight|nine))")
	if err != nil {
		log.Fatal(err)
	}

	maybeNumber := pattern.FindString(line)

	if len(maybeNumber) > 1 {
		return numberWords[maybeNumber]
	}
	return maybeNumber
}

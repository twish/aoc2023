package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func One(inputFile string) int {

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var total int = 0
	for _, line := range strings.Split(string(data), "\n") {
		nr, err := strconv.Atoi(getFirstNum(line) + getLastNum(line))
		if err != nil {
			log.Fatal(err)
		}
		total += nr
		//fmt.Println(total)
	}

	//fmt.Println("Total: ", total)
	return total
}

func getLastNum(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return string(line[i])
		}
	}
	log.Fatal("Could not find last number in ", line)
	return "0"
}

func getFirstNum(line string) string {
	for _, c := range line {
		if unicode.IsDigit(c) {
			return string(c)
		}
	}
	log.Fatal("Could not find first number in ", line)
	return "0"
}

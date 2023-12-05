package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type coordinate struct {
	line int
	col  int
}

type part struct {
	symbol           string
	partNumbers      []int
	symbolCoordinate coordinate
}

func One(inputFile string) int {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := make([][]rune, 0)
	for _, line := range strings.Split(string(data), "\n") {
		lines = append(lines, []rune(line))
	}

	parts, err := processRunes(lines)

	total := 0

	for _, p := range parts {
		for _, nr := range p.partNumbers {
			total += nr
		}
	}

	return total
}

func Two(inputFile string) int {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := make([][]rune, 0)
	for _, line := range strings.Split(string(data), "\n") {
		lines = append(lines, []rune(line))
	}

	parts, err := processRunes(lines)

	total := 0

	for _, p := range parts {
		if p.symbol == "*" && len(p.partNumbers) == 2 {
			total += p.partNumbers[0] * p.partNumbers[1]
		}
	}

	return total
}

func processRunes(lines [][]rune) (p []part, err error) {
	cols := len(lines[0])
	rows := len(lines)
	isSymbol, err := regexp.Compile("[^0-9.]")
	if err != nil {
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentRune := lines[i][j]
			if !isSymbol.MatchString(string(currentRune)) {
				continue
			}
			currentPart := part{
				symbol:           string(currentRune),
				symbolCoordinate: coordinate{line: i, col: j},
			}
			err = findPartNumbers(&currentPart, lines, cols, rows)
			if err != nil {
				return
			}
			p = append(p, currentPart)
		}
	}
	return
}

func findPartNumbers(p *part, lines [][]rune, cols int, rows int) (err error) {

	scanStart := coordinate{line: p.symbolCoordinate.line - 1, col: p.symbolCoordinate.col - 1}
	if scanStart.line < 0 {
		scanStart.line = 0
	}
	if scanStart.col < 0 {
		scanStart.col = 0
	}

	scanEnd := coordinate{line: p.symbolCoordinate.line + 1, col: p.symbolCoordinate.col + 1}
	if scanEnd.line > rows-1 {
		scanEnd.line = rows - 1
	}
	if scanEnd.col > cols-1 {
		scanEnd.col = cols - 1
	}

	for i := scanStart.line; i <= scanEnd.line; i++ {
		for j := scanStart.col; j <= scanEnd.col; j++ {
			if unicode.IsDigit(lines[i][j]) {
				head := j
				partNoRunes := make([]rune, 0)
				for unicode.IsDigit(lines[i][head]) {
					if head-1 < 0 {
						break
					}
					head--
				}
				if !unicode.IsDigit(lines[i][head]) {
					head++
				}
				for unicode.IsDigit(lines[i][head]) {
					partNoRunes = append(partNoRunes, lines[i][head])
					if head+1 >= cols {
						break
					}
					head++
				}

				partNo, err := strconv.Atoi(string(partNoRunes))
				if err != nil {
					return err
				}
				p.partNumbers = append(p.partNumbers, partNo)

				if head <= scanEnd.col {
					j = head
				} else {
					break
				}
			}
		}
	}
	return nil
}

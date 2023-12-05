package day4

import (
	"bufio"
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Card struct {
	id             int
	winningNumbers []int
	lottoNumbers   []int
}

func One(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := make([]Card, 0)
	for scanner.Scan() {
		card, err := parseToCard(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}

	total := 0
	for _, card := range cards {
		pow, err := card.Power()
		if err != nil {
			log.Fatal(err)
		}
		total += pow
	}

	return total
}

func Two(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := make([]Card, 0)
	for scanner.Scan() {
		card, err := parseToCard(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}

	return ProcessWins(cards, cards)

}

func ProcessWins(cards []Card, lookupCards []Card) (totalCards int) {
	for _, card := range cards {
		totalCards++
		newCards, err := card.GetWinnings(lookupCards)
		if err != nil {
			log.Fatal(err)
		}
		totalCards += ProcessWins(newCards, lookupCards)
	}
	return
}

func (c *Card) GetWinnings(allCards []Card) (winnings []Card, err error) {
	if len(allCards) <= 0 {
		err = errors.New("no cards to check winnings against")
		return
	}
	if len(c.lottoNumbers) <= 0 || len(c.winningNumbers) <= 0 {
		err = errors.New("winningNumbers or lottoNumbers was missing, cannot calculate power")
	}
	newWinCard := c.id
	for _, nr := range c.lottoNumbers {
		if slices.Contains(c.winningNumbers, nr) {
			newWinCard++
			if newWinCard < len(allCards) {
				winnings = append(winnings, allCards[newWinCard-1])
			}
		}
	}
	return
}

func (c *Card) Power() (pow int, err error) {
	if len(c.lottoNumbers) <= 0 || len(c.winningNumbers) <= 0 {
		err = errors.New("winningNumbers or lottoNumbers was missing, cannot calculate power")
	}
	for _, nr := range c.lottoNumbers {
		if slices.Contains(c.winningNumbers, nr) {
			if pow > 0 {
				pow *= 2
			} else {
				pow++
			}
		}
	}
	return
}

func parseToCard(text string) (card Card, err error) {
	parts := strings.Split(text, ":")
	if len(parts) != 2 {
		err = errors.New("splitting on : did not produce 2 parts")
		return
	}

	cardId, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Card ")))
	if err != nil {
		return
	}
	card.id = cardId

	numbers := strings.Split(parts[1], "|")
	if len(parts) != 2 {
		err = errors.New("splitting part[1] on | did not produce 2 parts")
		return
	}

	winningNumbers, err := readNumbers(numbers[0])
	if err != nil {
		return
	}
	card.winningNumbers = winningNumbers

	lottoNumbers, err := readNumbers(numbers[1])
	if err != nil {
		return
	}
	card.lottoNumbers = lottoNumbers

	return
}

func readNumbers(s string) (numbers []int, outErr error) {
	if len(s) <= 0 {
		outErr = errors.New("string was empty")
		return
	}
	runes := []rune(s)

	for i := 0; i < len(s); i++ {
		nrString := ""
		for unicode.IsDigit(runes[i]) {
			nrString += string(runes[i])

			if i == len(s)-1 {
				break
			}
			i++
		}
		if len(nrString) > 0 {
			nr, err := strconv.Atoi(nrString)
			if err != nil {
				outErr = err
				return
			}
			numbers = append(numbers, nr)
		}
	}
	return
}

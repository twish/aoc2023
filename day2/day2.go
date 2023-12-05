package day2

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets [][]int
}

type GameBag struct {
	Reds   int
	Greens int
	Blues  int
}

func One(inputFile string, bag GameBag) int {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGameRGB(strings.Split(string(data), "\n"))
	if err != nil {
		log.Fatal(err)
	}

	validGames := filterSlice(games, func(g game) bool {
		return g.Match(bag) != nil
	})

	total := 0

	for _, g := range validGames {
		total += g.id
	}

	return total
}

func Two(inputFile string) int {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGameRGB(strings.Split(string(data), "\n"))
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	for _, g := range games {
		total += g.Power()
	}

	return total
}

func filterSlice[T interface{ ~[]Ti }, Ti interface{}](stack T, filter func(Ti) bool) []Ti {
	var filteredSlice = make([]Ti, 0)
	for _, item := range stack {
		if filter(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}

func (g *game) Match(bag GameBag) *game {
	isMatch := true
	for _, set := range g.sets {
		isMatch = bag.Reds >= set[0] && bag.Greens >= set[1] && bag.Blues >= set[2]
		if !isMatch {
			return nil
		}
	}
	return g
}

func (g *game) Power() int {
	red, green, blue := 0, 0, 0
	for _, set := range g.sets {
		if red < set[0] {
			red = set[0]
		}
		if green < set[1] {
			green = set[1]
		}
		if blue < set[2] {
			blue = set[2]
		}
	}

	return red * green * blue
}

func parseGameRGB(lines []string) (res []game, err error) {
	for _, line := range lines {
		var game game = game{}
		lineParts := strings.Split(strings.TrimSpace(line), ":")
		if len(lineParts) != 2 {
			err = errors.New("line parts was not 2")
			return
		}
		game.id, err = strconv.Atoi(strings.Split(lineParts[0], " ")[1])
		if err != nil {
			return
		}

		game.sets, err = getRGB(lineParts[1])
		if err != nil {
			return
		}

		res = append(res, game)
	}
	return
}

func getRGB(rgbString string) (sets [][]int, err error) {
	var lookup = map[string]int{"red": 0, "green": 1, "blue": 2}
	for _, set := range strings.Split(rgbString, ";") {
		var rgbs = make([]int, 3)
		for _, rgb := range strings.Split(strings.TrimSpace(set), ",") {
			parts := strings.Split(strings.TrimSpace(rgb), " ")
			if len(parts) != 2 {
				err = errors.New("RGB part was not len 2")
				return
			}

			rgbs[lookup[parts[1]]], err = strconv.Atoi(parts[0])
			if err != nil {
				return
			}
		}
		sets = append(sets, rgbs)
	}
	return
}

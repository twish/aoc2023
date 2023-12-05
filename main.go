package main

import (
	"advent-of-code-2023/day1"
	"advent-of-code-2023/day2"
	"advent-of-code-2023/day3"
	"advent-of-code-2023/day4"
	"fmt"
)

func main() {

	// day1
	day1oneResult := day1.One("day1/input_data")
	fmt.Println("Day 1.1: ", day1oneResult)

	day1twoResult := day1.Two("day1/input_data")
	fmt.Println("Day 1.2: ", day1twoResult)

	// day2
	day2oneResult := day2.One("day2/input_data_1", day2.GameBag{12, 13, 14})
	fmt.Println("Day 2.1: ", day2oneResult)

	day2twoResult := day2.Two("day2/input_data_1")
	fmt.Println("Day 2.2: ", day2twoResult)

	// day3
	day3oneResult := day3.One("day3/input_data")
	fmt.Println("Day 3.1: ", day3oneResult)

	day3twoResult := day3.Two("day3/input_data")
	fmt.Println("Day 3.2: ", day3twoResult)

	//day4
	day4oneResult := day4.One("day4/input_data")
	fmt.Println("Day 4.1: ", day4oneResult)

	day4twoResult := day4.Two("day4/input_data")
	fmt.Println("Day 4.2: ", day4twoResult)
}

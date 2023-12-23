package day4

import (
	"fmt"
	"strings"

	"aoc_2023/utils/arrays"
	"aoc_2023/utils/math/intmath"
	"aoc_2023/utils/stringutils"
)

type Solver struct{}

type card struct {
	winning []int
	chosen  []int
}

func parseCard(line string) card {
	parts := strings.Split(line, ": ")
	parts = strings.Split(parts[1], " | ")
	winningNumbers := parseNumbers(parts[0])
	chosenNumbers := parseNumbers(parts[1])
	return card{
		winning: winningNumbers,
		chosen:  chosenNumbers,
	}
}

func parseNumbers(s string) []int {
	values := strings.Split(s, " ")
	values = arrays.Remove(values, stringutils.IsEmpty)
	numbers := arrays.Map(values, stringutils.Atoi)
	return numbers
}

func (c card) getMatchingNumbersCount() int {
	matchingNumbers := arrays.Intersection(c.winning, c.chosen)
	return len(matchingNumbers)
}

func (c card) getPoints() int {
	count := c.getMatchingNumbersCount()
	if count == 0 {
		return 0
	}
	return intmath.Power(2, count-1)
}

func getCards(input string) []card {
	lines := strings.Split(input, "\n")
	cards := arrays.Map(lines, parseCard)
	return cards
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	cards := getCards(input)
	points := arrays.Map(cards, card.getPoints)
	sum := arrays.Sum(points)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	cards := getCards(input)
	matchingNumbersCount := arrays.Map(cards, card.getMatchingNumbersCount)
	amounts := arrays.Map(cards, func(c card) int { return 1 })
	for i, count := range matchingNumbersCount {
		for j := 1; j <= count; j++ {
			amounts[i+j] += amounts[i]
		}
	}
	sum := arrays.Sum(amounts)
	return fmt.Sprintf("%d", sum)
}

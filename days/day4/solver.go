package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

type card struct {
	winning []int
	chosen  []int
}

func parseCard(line string) *card {
	parts := strings.Split(line, ": ")
	parts = strings.Split(parts[1], " | ")
	winningNumbers := parseNumbers(parts[0])
	chosenNumbers := parseNumbers(parts[1])
	return &card{
		winning: winningNumbers,
		chosen:  chosenNumbers,
	}
}

func parseNumbers(s string) []int {
	values := strings.Split(s, " ")
	numbers := make([]int, 0, len(values))
	for _, v := range values {
		if v == "" {
			continue
		}
		newNumber, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, newNumber)
	}
	return numbers
}

func (c *card) getMatches() []int {
	matches := make([]int, 0, len(c.winning))
	for _, w := range c.winning {
		if Contains(c.chosen, w) {
			matches = append(matches, w)
		}
	}
	return matches
}

func (c *card) getPoints() int {
	return int(math.Pow(2, float64(len(c.getMatches())-1)))
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func getCards(input string) []*card {
	lines := strings.Split(input, "\n")
	cards := make([]*card, 0, len(lines))
	for _, line := range lines {
		cards = append(cards, parseCard(line))
	}
	return cards
}

func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	cards := getCards(input)
	points := Map(cards, func(c *card) int { return c.getPoints() })
	sum := Sum(points)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}

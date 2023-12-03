package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type schematic [][]rune

func parseSchematic(input string) schematic {
	lines := strings.Split(input, "\n")
	s := make(schematic, len(lines))
	for i, line := range lines {
		s[i] = []rune(line)
	}
	return s
}

func (s *schematic) getPartNumbers() []int {
	partNumbers := make([]int, 0)
	for i := 0; i < len(*s); i++ {
		for j := 0; j < len((*s)[i]); j++ {
			if isSymbol((*s)[i][j]) {
				adjacents := s.getAdjacents(index{i, j})
				adjacents = s.removeDuplicates(index{i, j}, adjacents)
				for _, adjacent := range adjacents {
					if isDigit((*s)[adjacent.i][adjacent.j]) {
						partNumbers = append(partNumbers, s.getPartNumber(adjacent))
					}
				}
			}
		}
	}
	return partNumbers
}

func (s *schematic) getGearRatios() []int {
	partNumbers := make([]int, 0)
	for i := 0; i < len(*s); i++ {
		for j := 0; j < len((*s)[i]); j++ {
			if isPossibleGear((*s)[i][j]) {
				adjacents := s.getAdjacents(index{i, j})
				adjacents = s.removeDuplicates(index{i, j}, adjacents)
				parts := Filter(adjacents, func(index index) bool { return isDigit((*s)[index.i][index.j]) })
				newPartNumbers := Map(parts, func(index index) int { return s.getPartNumber(index) })
				newPartNumbers = Filter(newPartNumbers, func(number int) bool { return number != 0 })
				if len(newPartNumbers) == 2 {
					partNumbers = append(partNumbers, newPartNumbers[0]*newPartNumbers[1])
				}
			}
		}
	}
	return partNumbers
}

func (s *schematic) getAdjacents(i index) []index {
	adjacents := make([]index, 0, 8)
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if k == 0 && l == 0 {
				continue
			}
			if i.i+k < 0 || i.i+k >= len(*s) || i.j+l < 0 || i.j+l >= len((*s)[i.i+k]) {
				continue
			}
			adjacents = append(adjacents, index{i.i + k, i.j + l})
		}
	}
	return adjacents
}

func (s *schematic) removeDuplicates(i index, adjacents []index) []index {
	firstRow := Filter(adjacents, func(index index) bool { return index.i == i.i-1 })
	firstMiddle := Filter(firstRow, func(index index) bool { return index.j == i.j })
	if len(firstMiddle) != 0 {
		middleElement := firstMiddle[0]
		if isDigit((*s)[middleElement.i][middleElement.j]) {
			adjacents = Filter(adjacents, func(index index) bool { return index.i != middleElement.i || index.j == middleElement.j })
		} else {
			adjacents = Filter(adjacents, func(index index) bool { return index.i != middleElement.i || index.j != middleElement.j })
		}
	}
	lastRow := Filter(adjacents, func(index index) bool { return index.i == i.i+1 })
	lastMiddle := Filter(lastRow, func(index index) bool { return index.j == i.j })
	if len(lastMiddle) != 0 {
		middleElement := lastMiddle[0]
		if isDigit((*s)[middleElement.i][middleElement.j]) {
			adjacents = Filter(adjacents, func(index index) bool { return index.i != middleElement.i || index.j == middleElement.j })
		} else {
			adjacents = Filter(adjacents, func(index index) bool { return index.i != middleElement.i || index.j != middleElement.j })
		}
	}
	return adjacents
}

func Filter[T any](arr []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, i := range arr {
		if f(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func Map[T any, U any](arr []T, f func(T) U) []U {
	mapped := make([]U, 0)
	for _, i := range arr {
		mapped = append(mapped, f(i))
	}
	return mapped
}

func (s *schematic) getPartNumber(i index) int {
	first := i.j
	for ; first >= 0 && isDigit((*s)[i.i][first]); first-- {
	}
	first++
	digits := make([]rune, 0)
	for ; first < len((*s)[i.i]) && isDigit((*s)[i.i][first]); first++ {
		digits = append(digits, (*s)[i.i][first])
	}
	number, _ := strconv.Atoi(string(digits))
	return number
}

type index struct {
	i, j int
}

const symbols = "@#$%&*)-=+/"

func isSymbol(r rune) bool {
	return strings.Contains(symbols, string(r))
}

func isPossibleGear(r rune) bool {
	return r == '*'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func Sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	schematic := parseSchematic(input)
	partNumbers := schematic.getPartNumbers()
	sum := Sum(partNumbers)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	schematic := parseSchematic(input)
	gearRatios := schematic.getGearRatios()
	sum := Sum(gearRatios)
	return fmt.Sprintf("%d", sum)
}

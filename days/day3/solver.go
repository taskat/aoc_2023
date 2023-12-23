package day3

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/math/intmath"
	"aoc_2023/utils/stringutils"
	"fmt"
	"strings"
)

type Solver struct{}

const symbols = "@#$%&*)-=+/"

type schematic [][]rune

func parseSchematic(input string) schematic {
	lines := strings.Split(input, "\n")
	s := arrays.Map(lines, func(line string) []rune { return []rune(line) })
	return s
}

func (s *schematic) isDigit(idx index) bool {
	return stringutils.IsDigit((*s)[idx.i][idx.j])
}

func (s *schematic) isSymbol(idx index) bool {
	return strings.Contains(symbols, string((*s)[idx.i][idx.j]))
}

func (s *schematic) isPossibleGear(idx index) bool {
	return (*s)[idx.i][idx.j] == '*'
}

func (s *schematic) getPartNumberByIdx(idx index) []int {
	if !s.isSymbol(idx) {
		return make([]int, 0)
	}
	adjacents := s.getAdjacentIndices(idx)
	adjacents = s.removeDuplicates(idx, adjacents)
	newNumberIndices := arrays.Filter(adjacents, s.isDigit)
	return arrays.Map(newNumberIndices, s.getPartNumber)
}

func (s *schematic) getGearRatiosByIdx(idx index) []int {
	if !s.isPossibleGear(idx) {
		return make([]int, 0)
	}
	adjacents := s.getAdjacentIndices(idx)
	adjacents = s.removeDuplicates(idx, adjacents)
	parts := arrays.Filter(adjacents, s.isDigit)
	newPartNumbers := arrays.Map(parts, s.getPartNumber)
	newPartNumbers = arrays.Filter(newPartNumbers, intmath.IsNotZero[int])
	if len(newPartNumbers) == 2 {
		return []int{newPartNumbers[0] * newPartNumbers[1]}
	}
	return make([]int, 0)
}

func (s *schematic) getParts(listParts func(index) []int) []int {
	partNumbers := make([]int, 0)
	for i := 0; i < len(*s); i++ {
		for j := 0; j < len((*s)[i]); j++ {
			idx := index{i, j}
			partNumbers = append(partNumbers, listParts(idx)...)
		}
	}
	return partNumbers
}

func (s *schematic) getAdjacentIndices(i index) []index {
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

func (s *schematic) removeDuplicates(idx index, adjacents []index) []index {
	middleElems := arrays.Filter(adjacents, idx.isVerticalNeighbor)
	for _, middleElem := range middleElems {
		var remover func(index) bool
		if s.isDigit(middleElem) {
			remover = func(index index) bool { return index.inSameRow(middleElem) && !index.same(middleElem) }
		} else {
			remover = idx.same
		}
		adjacents = arrays.Remove(adjacents, remover)
	}
	return adjacents
}

func (s *schematic) getNumberStartIndex(base, row int) int {
	first := base
	for ; first >= 0 && s.isDigit(index{i: row, j: first}); first-- {
	}
	first++
	return first
}

func (s *schematic) getNumberEndIndex(base, row int) int {
	last := base
	for ; last < len((*s)[row]) && s.isDigit(index{i: row, j: last}); last++ {
	}
	last--
	return last
}

func (s *schematic) getPartNumber(i index) int {
	first := s.getNumberStartIndex(i.j, i.i)
	last := s.getNumberEndIndex(i.j, i.i)
	number := stringutils.Atoi(string((*s)[i.i][first : last+1]))
	return number
}

type index struct {
	i, j int
}

func (idx *index) inSameColumn(other index) bool {
	return idx.j == other.j
}

func (idx *index) inSameRow(other index) bool {
	return idx.i == other.i
}

func (idx *index) isVerticalNeighbor(other index) bool {
	return idx.j == other.j && intmath.Abs(idx.i-other.i) == 1
}

func (idx *index) isHorizontalNeighbor(other index) bool {
	return idx.i == other.i && intmath.Abs(idx.j-other.j) == 1
}

func (idx *index) same(other index) bool {
	return idx.i == other.i && idx.j == other.j
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	schematic := parseSchematic(input)
	partNumbers := schematic.getParts(schematic.getPartNumberByIdx)
	sum := arrays.Sum(partNumbers)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	schematic := parseSchematic(input)
	gearRatios := schematic.getParts(schematic.getGearRatiosByIdx)
	sum := arrays.Sum(gearRatios)
	return fmt.Sprintf("%d", sum)
}

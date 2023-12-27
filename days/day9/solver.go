package day9

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/math/intmath"
	"aoc_2023/utils/stringutils"
	"fmt"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	sequences := parseSequences(lines)
	predictions := arrays.Map(sequences, sequence.predict)
	sum := arrays.Sum(predictions)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	sequences := parseSequences(lines)
	predictions := arrays.Map(sequences, sequence.predictPast)
	sum := arrays.Sum(predictions)
	return fmt.Sprintf("%d", sum)
}

type sequence []int

func parseSequence(input string) sequence {
	parts := strings.Split(input, " ")
	return sequence(arrays.Map(parts, stringutils.Atoi))
}

func (s sequence) pastPredictor(idx, value int) int {
	if idx == 0 {
		return value
	}
	return value - s[idx-1]
}

func (s sequence) predictor(idx, value int) int {
	if idx == len(s)-1 {
		return value
	}
	return s[idx+1] - value
}

func (s sequence) predict() int {
	if arrays.All(s, intmath.IsZero) {
		return 0
	}
	var diff sequence = arrays.Map_i(s, s.predictor)
	diff = diff[:len(diff)-1]
	diff = append(diff, diff.predict())
	return diff[len(diff)-1] + s[len(s)-1]
}

func (s sequence) predictPast() int {
	if arrays.All(s, intmath.IsZero) {
		return 0
	}
	var diff sequence = arrays.Map_i(s, s.pastPredictor)
	diff = diff[1:]
	diff = append(sequence{diff.predictPast()}, diff...)
	return s[0] - diff[0]
}

func parseSequences(lines []string) []sequence {
	return arrays.Map(lines, parseSequence)
}

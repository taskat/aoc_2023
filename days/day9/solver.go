package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type sequence []int

func parseSequence(input string) sequence {
	parts := strings.Split(input, " ")
	return sequence(Map(parts, atoi))
}

func (s sequence) predict() int {
	if All(s, func(n int) bool { return n == 0 }) {
		return 0
	}
	diff := make(sequence, 0, len(s))
	for i := 1; i < len(s); i++ {
		diff = append(diff, s[i]-s[i-1])
	}
	diff = append(diff, diff.predict())
	return diff[len(diff)-1] + s[len(s)-1]
}

func (s sequence) predictPast() int {
	if All(s, func(n int) bool { return n == 0 }) {
		return 0
	}
	diff := make(sequence, 0, len(s))
	for i := 1; i < len(s); i++ {
		diff = append(diff, s[i]-s[i-1])
	}
	diff = append(sequence{diff.predictPast()}, diff...)
	return s[0] - diff[0]
}

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Map[T, U any](arr []T, mapper func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = mapper(v)
	}
	return result
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseSequences(input string) []sequence {
	parts := strings.Split(input, "\n")
	return Map(parts, parseSequence)
}

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	sequences := parseSequences(input)
	predictions := Map(sequences, sequence.predict)
	sum := Sum(predictions)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	sequences := parseSequences(input)
	predictions := Map(sequences, sequence.predictPast)
	sum := Sum(predictions)
	return fmt.Sprintf("%d", sum)
}

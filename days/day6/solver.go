package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type race struct {
	time   int
	record int
}

func parseRaces(input string) []race {
	lines := strings.Split(input, "\n")
	timeLine := strings.TrimSpace(strings.ReplaceAll(lines[0], "Time:", ""))
	recordLine := strings.TrimSpace(strings.ReplaceAll(lines[1], "Distance:", ""))
	timeStrings := strings.Split(timeLine, " ")
	timeStrings = Filter(timeStrings, func(s string) bool { return s != "" })
	times := Map(timeStrings, func(s string) int { return atoi(s) })
	recordStrings := strings.Split(recordLine, " ")
	recordStrings = Filter(recordStrings, func(s string) bool { return s != "" })
	records := Map(recordStrings, func(s string) int { return atoi(s) })
	races := make([]race, len(timeStrings))
	for i := 0; i < len(timeStrings); i++ {
		races[i] = race{time: times[i], record: records[i]}
	}
	return races
}

func (r race) distances() []int {
	distances := make([]int, r.time)
	for i := 1; i < r.time; i++ {
		distances[i] = i * (r.time - i)
	}
	return distances
}

func (r race) waysToWin() int {
	counter := 0
	for _, distance := range r.distances() {
		if distance > r.record {
			counter++
		}
	}
	return counter
}

func atoi(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

func Filter[T any](s []T, fn func(T) bool) []T {
	var p []T // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func Map[T, U any](s []T, fn func(T) U) []U {
	p := make([]U, len(s))
	for i, v := range s {
		p[i] = fn(v)
	}
	return p
}

func Product(slice []int) int {
	product := 1
	for _, elem := range slice {
		product *= elem
	}
	return product
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	races := parseRaces(input)
	waysToWin := Map(races, race.waysToWin)
	product := Product(waysToWin)
	return fmt.Sprintf("%d", product)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}

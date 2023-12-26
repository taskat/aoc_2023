package day6

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/math/intmath"
	"aoc_2023/utils/stringutils"
	"aoc_2023/utils/types"
	"fmt"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	races := parseRaces(lines)
	waysToWin := arrays.Map(races, race.waysToWin)
	product := arrays.Product(waysToWin)
	return fmt.Sprintf("%d", product)
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	race := parseOneRace(lines)
	waysToWin := race.waysToWin()
	return fmt.Sprintf("%d", waysToWin)
}

type race struct {
	time   int
	record int
}

func parseRaces(lines []string) []race {
	timeLine := strings.TrimSpace(strings.ReplaceAll(lines[0], "Time:", ""))
	recordLine := strings.TrimSpace(strings.ReplaceAll(lines[1], "Distance:", ""))
	timeStrings := strings.Split(timeLine, " ")
	timeStrings = arrays.Filter(timeStrings, stringutils.IsNotEmpty)
	times := arrays.Map(timeStrings, stringutils.Atoi)
	recordStrings := strings.Split(recordLine, " ")
	recordStrings = arrays.Filter(recordStrings, stringutils.IsNotEmpty)
	records := arrays.Map(recordStrings, stringutils.Atoi)
	pairs := arrays.Pair(times, records)
	races := arrays.Map(pairs, func(p types.Pair[int, int]) race { return race{time: p.First, record: p.Second} })
	return races
}

func (r race) distances() []int {
	distances := arrays.Map_i(make([]int, r.time), func(i int, _ int) int { return i * (r.time - i) })
	return distances
}

func (r race) waysToWin() int {
	x1, x2 := intmath.SolveQuadratic(1, -r.time, r.record)
	from := intmath.Ceil[int](x1)
	to := intmath.Floor[int](x2)
	return intmath.Abs(to - from + 1)
}

func parseOneRace(lines []string) race {
	timeLine := strings.TrimSpace(strings.ReplaceAll(lines[0], "Time:", ""))
	recordLine := strings.TrimSpace(strings.ReplaceAll(lines[1], "Distance:", ""))
	timeString := strings.ReplaceAll(timeLine, " ", "")
	time := stringutils.Atoi(timeString)
	recordString := strings.ReplaceAll(recordLine, " ", "")
	record := stringutils.Atoi(recordString)
	return race{time: time, record: record}
}

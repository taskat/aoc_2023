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

type race struct {
	time   int
	record int
}

func parseRaces(input string) []race {
	lines := strings.Split(input, "\n")
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
	distances := arrays.Mapi(make([]int, r.time), func(i int, _ int) int { return i * (r.time - i) })
	return distances
}

func (r race) waysToWin() int {
	x1, x2 := intmath.SolveQuadratic(1, -r.time, r.record)
	from := intmath.Ceil[int](x1)
	to := intmath.Floor[int](x2)
	return intmath.Abs(to - from + 1)
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	races := parseRaces(input)
	waysToWin := arrays.Map(races, race.waysToWin)
	product := arrays.Product(waysToWin)
	return fmt.Sprintf("%d", product)
}

func parseoneRace(input string) race {
	lines := strings.Split(input, "\n")
	timeLine := strings.TrimSpace(strings.ReplaceAll(lines[0], "Time:", ""))
	recordLine := strings.TrimSpace(strings.ReplaceAll(lines[1], "Distance:", ""))
	timeString := strings.ReplaceAll(timeLine, " ", "")
	time := stringutils.Atoi(timeString)
	recordString := strings.ReplaceAll(recordLine, " ", "")
	record := stringutils.Atoi(recordString)
	return race{time: time, record: record}
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	race := parseoneRace(input)
	waysToWin := race.waysToWin()
	return fmt.Sprintf("%d", waysToWin)
}

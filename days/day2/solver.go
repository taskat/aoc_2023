package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type game struct {
	id    int
	cubes []revealed
}

func newGame(line string) game {
	line = strings.ReplaceAll(line, "Game ", "")
	parts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(parts[0])
	turns := strings.Split(parts[1], "; ")
	cubes := make([]revealed, len(turns))
	for i, turn := range turns {
		cubes[i] = newRevealed(turn)
	}
	return game{id, cubes}
}

func (g *game) isPossible(l limit) bool {
	for _, cube := range g.cubes {
		for color, count := range cube {
			if count > l[color] {
				return false
			}
		}
	}
	return true
}

type revealed map[string]int

func newRevealed(line string) revealed {
	turns := strings.Split(line, ", ")
	revealed := make(revealed, len(turns))
	for _, turn := range turns {
		parts := strings.Split(turn, " ")
		revealed[parts[1]], _ = strconv.Atoi(parts[0])
	}
	return revealed
}

func parseGames(input string) []game {
	lines := strings.Split(input, "\n")
	games := make([]game, len(lines))
	for i, line := range lines {
		games[i] = newGame(line)
	}
	return games
}

type limit map[string]int

func getLimit(extraParams []any) limit {
	l := make(limit, 3)
	l["red"], _ = strconv.Atoi(extraParams[0].(string))
	l["green"], _ = strconv.Atoi(extraParams[1].(string))
	l["blue"], _ = strconv.Atoi(extraParams[2].(string))
	return l
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

func Sum[T ~int | ~float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	limit := getLimit(extraParams)
	games := parseGames(input)
	possibleGames := Filter(games, func(g game) bool { return g.isPossible(limit) })
	possibleGameIds := Map(possibleGames, func(g game) int { return g.id })
	sum := Sum(possibleGameIds)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}

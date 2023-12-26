package day2

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/maps"
	"aoc_2023/utils/stringutils"
	"fmt"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	limit := getLimit(extraParams)
	games := parseGames(lines)
	possibleGames := arrays.Filter(games, func(g game) bool { return g.isPossible(limit) })
	possibleGameIds := arrays.Map(possibleGames, func(g game) int { return g.id })
	sum := arrays.Sum(possibleGameIds)
	return fmt.Sprintf("%d", sum)
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	games := parseGames(lines)
	powers := arrays.Map(games, func(g game) int { return g.powerOfMinimal() })
	sum := arrays.Sum(powers)
	return fmt.Sprintf("%d", sum)
}

type game struct {
	id     int
	rounds []round
}

func newGame(line string) game {
	line = strings.ReplaceAll(line, "Game ", "")
	parts := strings.Split(line, ": ")
	id := stringutils.Atoi(parts[0])
	turns := strings.Split(parts[1], "; ")
	cubes := arrays.Map(turns, parseRound)
	return game{id, cubes}
}

func (g *game) isPossible(limit round) bool {
	return arrays.All(g.rounds, func(r round) bool {
		return r.isPossible(limit)
	})
}

func (g *game) minimalNecessaryCubes() round {
	minimal := g.rounds[0]
	arrays.ForEach(g.rounds, func(r round) { r.updateMinimal(&minimal) })
	return minimal
}

func (g *game) powerOfMinimal() int {
	minimal := g.minimalNecessaryCubes()
	minimalCubes := maps.MapToArray(minimal, func(_ string, count int) int { return count })
	return arrays.Product(minimalCubes)
}

type round map[string]int

func parseTurn(turn string) (string, int) {
	parts := strings.Split(turn, " ")
	return parts[1], stringutils.Atoi(parts[0])
}

func parseRound(line string) round {
	turns := strings.Split(line, ", ")
	revealed := arrays.MapToMap(turns, parseTurn)
	return revealed
}

func (r round) isPossible(limit round) bool {
	return maps.All(r, func(color string, count int) bool { return count <= limit[color] })
}

func (r round) updateMinimal(minimal *round) {
	maps.ForEach(r, func(color string, count int) {
		if count > (*minimal)[color] {
			(*minimal)[color] = count
		}
	})
}

func parseGames(lines []string) []game {
	games := arrays.Map(lines, newGame)
	return games
}

func getLimit(extraParams []any) round {
	l := make(round, 3)
	l["red"] = stringutils.Atoi(extraParams[0].(string))
	l["green"] = stringutils.Atoi(extraParams[1].(string))
	l["blue"] = stringutils.Atoi(extraParams[2].(string))
	return l
}

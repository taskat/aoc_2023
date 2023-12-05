package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type smartRange struct {
	src    int
	dst    int
	length int
}

func parseSmartRange(s string) smartRange {
	parts := strings.Split(s, " ")
	dst, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	src, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	length, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	return smartRange{
		src:    src,
		dst:    dst,
		length: length,
	}
}

func (sm *smartRange) contains(value int) bool {
	return value >= sm.src && value <= sm.src+sm.length
}

func (sm *smartRange) getMappedValue(value int) int {
	if !sm.contains(value) {
		return value
	}
	return sm.dst + value - sm.src
}

type gardenMapping struct {
	from   string
	to     string
	ranges []smartRange
}

func parseGardenMapping(input string) gardenMapping {
	lines := strings.Split(input, "\n")
	title := lines[0]
	title = strings.ReplaceAll(title, " map:", "")
	parts := strings.Split(title, "-")
	from := parts[0]
	to := parts[2]
	ranges := make([]smartRange, len(lines)-1)
	for _, line := range lines[1:] {
		newRange := parseSmartRange(line)
		ranges = append(ranges, newRange)
	}
	return gardenMapping{
		from:   from,
		to:     to,
		ranges: ranges,
	}
}

func (gm *gardenMapping) mapValue(value int) int {
	for _, r := range gm.ranges {
		if r.contains(value) {
			return r.getMappedValue(value)
		}
	}
	return value
}

func parseSeeds(line string) []int {
	line = strings.ReplaceAll(line, "seeds: ", "")
	parts := strings.Split(line, " ")
	seeds := make([]int, len(parts))
	for i, part := range parts {
		seeds[i], _ = strconv.Atoi(part)
	}
	return seeds
}

func parseMappings(input string) ([]int, map[string]gardenMapping) {
	lines := strings.Split(input, "\n\n")
	seeds := parseSeeds(lines[0])
	mappings := make([]gardenMapping, len(lines)-1)
	for i, line := range lines[1:] {
		mappings[i] = parseGardenMapping(line)
	}
	gardens := make(map[string]gardenMapping, len(mappings))
	for _, mapping := range mappings {
		gardens[mapping.from] = mapping
	}
	return seeds, gardens
}

func findLocations(seeds []int, gardens map[string]gardenMapping) []int {
	start := "seed"
	goal := "location"
	current := start
	values := make([]int, len(seeds))
	copy(values, seeds)
	for current != goal {
		garden := gardens[current]
		for i, value := range values {
			values[i] = garden.mapValue(value)
		}
		current = garden.to
	}
	return values
}

func Min(values []int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	seeds, mappings := parseMappings(input)
	locations := findLocations(seeds, mappings)
	minLocation := Min(locations)
	return fmt.Sprintf("%d", minLocation)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}

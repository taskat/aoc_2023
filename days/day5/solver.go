package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

type Range struct {
	from int
	to   int
}

func newRange(from, to int) Range {
	if to < from {
		panic(fmt.Sprintf("Invalid range: %d-%d", from, to))
	}
	return Range{
		from: from,
		to:   to,
	}
}

func newRangeWithLength(from, length int) Range {
	return Range{
		from: from,
		to:   from + length,
	}
}

func (r *Range) contains(value int) bool {
	return value >= r.from && value < r.to
}

func (r *Range) containsRange(other Range) bool {
	return r.contains(other.from) && r.contains(other.to-1)
}

func (r *Range) hasIntersection(other Range) bool {
	return r.contains(other.from) || r.contains(other.to-1) || other.contains(r.from) || other.contains(r.to-1)
}

func (r *Range) length() int {
	return r.to - r.from
}

func (r *Range) String() string {
	return fmt.Sprintf("From: %d, To: %d", r.from, r.to)
}

type mapping struct {
	Range
	dstStart int
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseMapping(s string) mapping {
	parts := strings.Split(s, " ")
	dstStart := atoi(parts[0])
	srcStart := atoi(parts[1])
	srcLength := atoi(parts[2])
	return mapping{
		Range:    newRangeWithLength(srcStart, srcLength),
		dstStart: dstStart,
	}
}

func (m *mapping) getMappedValue(value int) int {
	if !m.contains(value) {
		fmt.Println("should not happen in getMappedValue")
		return value
	}
	return m.dstStart + value - m.from
}

func (m *mapping) String() string {
	return fmt.Sprintf("Range: %v, DstStart: %d", m.Range.String(), m.dstStart)
}

func (m *mapping) mapRange(r Range) (Range, []Range) {
	switch {
	case m.containsRange(r):
		return newRangeWithLength(m.getMappedValue(r.from), r.length()), []Range{}
	case !m.hasIntersection(r):
		fmt.Println("should not happen")
		return Range{}, []Range{r}
	case m.from > r.from && m.to < r.to:
		return newRange(m.getMappedValue(m.from), m.getMappedValue(m.to-1)+1), []Range{
			newRange(r.from, m.from),
			newRange(m.to, r.to),
		}
	case m.from <= r.from && m.to < r.to:
		return newRange(m.getMappedValue(r.from), m.getMappedValue(m.to-1)+1), []Range{
			newRange(m.to, r.to),
		}
	case m.from >= r.from && m.to >= r.to:
		return newRange(m.getMappedValue(m.from), m.getMappedValue(r.to-1)+1), []Range{
			newRange(r.from, m.from),
		}
	default:
		fmt.Println(m)
		fmt.Println(r.String())
		panic("should not happen")
	}
}

type gardenMapping struct {
	from     string
	to       string
	mappings []mapping
}

func parseGardenMapping(input string) gardenMapping {
	lines := strings.Split(input, "\n")
	title := lines[0]
	title = strings.ReplaceAll(title, " map:", "")
	parts := strings.Split(title, "-")
	from := parts[0]
	to := parts[2]
	mappings := make([]mapping, 0, len(lines)-1)
	for _, line := range lines[1:] {
		newMapping := parseMapping(line)
		mappings = append(mappings, newMapping)
	}
	return gardenMapping{
		from:     from,
		to:       to,
		mappings: mappings,
	}
}

func (gm *gardenMapping) mapValue(value int) int {
	for _, m := range gm.mappings {
		if m.contains(value) {
			return m.getMappedValue(value)
		}
	}
	return value
}

func (gm *gardenMapping) mapRanges(ranges []Range) []Range {
	newRanges := make([]Range, 0, len(ranges))
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]
		wasIntersected := false
		for _, m := range gm.mappings {
			if !m.hasIntersection(r) {
				continue
			}
			wasIntersected = true
			newRange, rem := m.mapRange(r)
			if newRange != (Range{}) {
				newRanges = append(newRanges, newRange)
			}
			ranges = append(ranges, rem...)
			break
		}
		if !wasIntersected {
			newRanges = append(newRanges, r)
		}
	}
	return newRanges
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

func findLocation(seeds int, gardens map[string]gardenMapping) int {
	start := "seed"
	goal := "location"
	current := start
	value := seeds
	for current != goal {
		garden := gardens[current]
		value = garden.mapValue(value)
		current = garden.to
	}
	return value
}

func findLocations(seeds []int, gardens map[string]gardenMapping) []int {
	locations := make([]int, len(seeds))
	for i, seed := range seeds {
		locations[i] = findLocation(seed, gardens)
	}
	return locations
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

func parseSeedRanges(line string) []Range {
	line = strings.ReplaceAll(line, "seeds: ", "")
	parts := strings.Split(line, " ")
	ranges := make([]Range, 0, len(parts)/2)
	for i := 0; i < len(parts); i += 2 {
		from := atoi(parts[i])
		length := atoi(parts[i+1])
		ranges = append(ranges, newRangeWithLength(from, length))
	}
	return ranges
}

func parseMappings_2(input string) ([]Range, map[string]gardenMapping) {
	blocks := strings.Split(input, "\n\n")
	seeds := parseSeedRanges(blocks[0])
	mappings := make([]gardenMapping, len(blocks)-1)
	for i, block := range blocks[1:] {
		mappings[i] = parseGardenMapping(block)
	}
	gardens := make(map[string]gardenMapping, len(mappings))
	for _, mapping := range mappings {
		gardens[mapping.from] = mapping
	}
	return seeds, gardens
}

func getMin(ranges []Range) int {
	min := ranges[0].from
	for _, r := range ranges {
		if r.from < min {
			min = r.from
		}
	}
	return min
}

func findMinLocation(r Range, gardens map[string]gardenMapping) int {
	start := "seed"
	goal := "location"
	current := start
	ranges := make([]Range, 1)
	ranges[0] = r
	for current != goal {
		garden := gardens[current]
		ranges = garden.mapRanges(ranges)
		current = garden.to
	}
	return getMin(ranges)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	seeds, mappings := parseMappings_2(input)
	min := findMinLocation(seeds[0], mappings)
	for _, seedRange := range seeds {
		location := findMinLocation(seedRange, mappings)
		if location < min {
			min = location
		}
	}
	return fmt.Sprintf("%d", min)
}

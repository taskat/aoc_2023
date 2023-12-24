package day5

import (
	"aoc_2023/utils/arrays"
	ru "aoc_2023/utils/rangeutils"
	"aoc_2023/utils/stringutils"
	"fmt"
	"strings"
)

type Solver struct{}

type mapping struct {
	ru.Range
	dstStart int
}

func parseMapping(s string) mapping {
	parts := strings.Split(s, " ")
	dstStart := stringutils.Atoi(parts[0])
	srcStart := stringutils.Atoi(parts[1])
	srcLength := stringutils.Atoi(parts[2])
	return mapping{
		Range:    *ru.NewRangeWithLength(srcStart, srcLength),
		dstStart: dstStart,
	}
}

func (m mapping) diff() int {
	return m.dstStart - m.From()
}

func (m *mapping) mapValue(value int) int {
	if !m.Contains(value) {
		return value
	}
	return value + m.diff()
}

func (m *mapping) String() string {
	return fmt.Sprintf("Range: %v, DstStart: %d", m.Range.String(), m.dstStart)
}

func (m *mapping) mapRange(r ru.Range) (*ru.Range, []ru.Range) {
	switch {
	case m.ContainsRange(r):
		return ru.NewRangeWithLength(m.mapValue(r.From()), r.Length()), []ru.Range{}
	case !m.HasIntersection(r):
		return nil, []ru.Range{r}
	case m.From() > r.From() && m.To() < r.To():
		return ru.NewRange(m.mapValue(m.From()), m.mapValue(m.To()-1)+1), []ru.Range{
			*ru.NewRange(r.From(), m.From()),
			*ru.NewRange(m.To(), r.To()),
		}
	case m.From() <= r.From() && m.To() < r.To():
		return ru.NewRange(m.mapValue(r.From()), m.mapValue(m.To()-1)+1), []ru.Range{
			*ru.NewRange(m.To(), r.To()),
		}
	case m.From() >= r.From() && m.To() >= r.To():
		return ru.NewRange(m.mapValue(m.From()), m.mapValue(r.To()-1)+1), []ru.Range{
			*ru.NewRange(r.From(), m.From()),
		}
	default:
		panic("ranges has no relation")
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
	mappings := arrays.Map(lines[1:], parseMapping)
	return gardenMapping{
		from:     from,
		to:       to,
		mappings: mappings,
	}
}

func (gm gardenMapping) keyValue() (string, gardenMapping) {
	return gm.from, gm
}

func (gm *gardenMapping) mapValue(value int) int {
	mapping, ok := arrays.Find(gm.mappings, func(m mapping) bool { return m.Contains(value) })
	if !ok {
		return value
	}
	return mapping.mapValue(value)
}

func (gm *gardenMapping) mapRanges(ranges []ru.Range) []ru.Range {
	newRanges := make([]ru.Range, 0, len(ranges))
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]
		wasIntersected := false
		for _, m := range gm.mappings {
			if !m.HasIntersection(r) {
				continue
			}
			wasIntersected = true
			mapped, rem := m.mapRange(r)
			if mapped != nil {
				newRanges = append(newRanges, *mapped)
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
	seeds := arrays.Map(parts, stringutils.Atoi)
	return seeds
}

func parseMappings(input string) ([]int, map[string]gardenMapping) {
	lines := strings.Split(input, "\n\n")
	seeds := parseSeeds(lines[0])
	mappings := arrays.Map(lines[1:], parseGardenMapping)
	gardens := arrays.MapToMap(mappings, gardenMapping.keyValue)
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
	return arrays.Map(seeds, func(seed int) int { return findLocation(seed, gardens) })
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	seeds, mappings := parseMappings(input)
	locations := findLocations(seeds, mappings)
	minLocation := arrays.Min(locations)
	return fmt.Sprintf("%d", minLocation)
}

func parseSeedRanges(line string) []ru.Range {
	line = strings.ReplaceAll(line, "seeds: ", "")
	parts := strings.Split(line, " ")
	ranges := make([]ru.Range, 0, len(parts)/2)
	for i := 0; i < len(parts); i += 2 {
		from := stringutils.Atoi(parts[i])
		length := stringutils.Atoi(parts[i+1])
		ranges = append(ranges, *ru.NewRangeWithLength(from, length))
	}
	return ranges
}

func parseMappings_2(input string) ([]ru.Range, map[string]gardenMapping) {
	blocks := strings.Split(input, "\n\n")
	seeds := parseSeedRanges(blocks[0])
	mappings := arrays.Map(blocks[1:], parseGardenMapping)
	gardens := arrays.MapToMap(mappings, gardenMapping.keyValue)
	return seeds, gardens
}

func findMinLocation(r ru.Range, gardens map[string]gardenMapping) int {
	start := "seed"
	goal := "location"
	current := start
	ranges := make([]ru.Range, 1)
	ranges[0] = r
	for current != goal {
		garden := gardens[current]
		ranges = garden.mapRanges(ranges)
		current = garden.to
	}
	starts := arrays.Map(ranges, ru.Range.From)
	return arrays.Min(starts)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	seeds, mappings := parseMappings_2(input)
	locations := arrays.Map(seeds, func(seed ru.Range) int { return findMinLocation(seed, mappings) })
	min := arrays.Min(locations)
	return fmt.Sprintf("%d", min)
}

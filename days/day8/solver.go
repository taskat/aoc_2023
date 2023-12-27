package day8

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/maps"
	"aoc_2023/utils/math/intmath"
	"fmt"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	directions, nodes := parseInput(lines)
	_, start, _ := maps.Find(nodes, func(_ string, n *node) bool { return n.isStart() })
	steps := start.getCycleLength(directions, nodes, (*node).isGoal)
	return fmt.Sprintf("%d", steps)
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	directions, nodes := parseInput(lines)
	starts := maps.Filter(nodes, func(_ string, n *node) bool { return n.isGhostStart() })
	cycles := maps.MapToArray(starts, func(_ string, n *node) int { return n.getCycleLength(directions, nodes, (*node).isGhostGoal) })
	minSteps := getMinStep(cycles)
	return fmt.Sprintf("%d", minSteps)
}

type dir rune

const (
	left  dir = 'L'
	right dir = 'R'
)

func parseDirections(line string) []dir {
	return arrays.Map([]rune(line), func(c rune) dir { return dir(c) })
}

type node struct {
	label     string
	neighbors map[dir]string
}

func (n *node) isGhostStart() bool {
	return n.label[len(n.label)-1] == 'A'
}

func (n *node) isGhostGoal() bool {
	return n.label[len(n.label)-1] == 'Z'
}

func (n *node) isStart() bool {
	return n.label == "AAA"
}

func (n *node) isGoal() bool {
	return n.label == "ZZZ"
}

func (n *node) getCycleLength(directions []dir, nodes map[string]*node, isGoal func(n *node) bool) int {
	current := n
	counter := 0
	for !isGoal(current) {
		for i, d := range directions {
			if isGoal(current) {
				return counter + i
			}
			current = nodes[current.neighbors[d]]
		}
		counter += len(directions)
	}
	return counter
}

func parseNode(line string) *node {
	// parts := strings.Split(line, " = ")
	// label := parts[0]
	label := line[0:3]
	// neighborsString := strings.Trim(parts[1], "()")
	// neighborsList := strings.Split(neighborsString, ", ")
	neighbors := map[dir]string{
		// left:  neighborsList[0],
		// right: neighborsList[1],
		left:  line[7:10],
		right: line[12:15],
	}
	return &node{label, neighbors}
}

func parseInput(lines []string) ([]dir, map[string]*node) {
	directions := parseDirections(lines[0])
	nodes := arrays.MapToMap(lines[2:], func(line string) (string, *node) {
		n := parseNode(line)
		return n.label, n
	})
	return directions, nodes
}

func getMinStep(cycles []int) int {
	minSteps := intmath.Lcm(cycles[0], cycles[1])
	for i := 2; i < len(cycles); i++ {
		minSteps = intmath.Lcm(minSteps, cycles[i])
	}
	return minSteps
}

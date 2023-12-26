package day8

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/maps"
	"aoc_2023/utils/math/intmath"
	"fmt"
	"strings"
)

type Solver struct{}

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

func (n *node) isStart() bool {
	return n.label[len(n.label)-1] == 'A'
}

func (n *node) isGoal() bool {
	return n.label[len(n.label)-1] == 'Z'
}

func (n *node) getCycleLength(directions []dir, nodes map[string]*node) int {
	current := n
	counter := 0
	for !current.isGoal() {
		for i, d := range directions {
			if current.isGoal() {
				return counter + i
			}
			current = nodes[current.neighbors[d]]
		}
		counter += len(directions)
	}
	return counter
}

func parseNode(line string) *node {
	parts := strings.Split(line, " = ")
	label := parts[0]
	neighborsString := strings.Trim(parts[1], "()")
	neighborsList := strings.Split(neighborsString, ", ")
	neighbors := map[dir]string{
		left:  neighborsList[0],
		right: neighborsList[1],
	}
	return &node{label, neighbors}
}

func parseInput(input string) ([]dir, map[string]*node) {
	parts := strings.Split(input, "\n\n")
	directions := parseDirections(parts[0])
	lines := strings.Split(parts[1], "\n")
	nodes := arrays.MapToMap(lines, func(line string) (string, *node) {
		n := parseNode(line)
		return n.label, n
	})
	return directions, nodes
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	directions, nodes := parseInput(input)
	start := "AAA"
	goal := "ZZZ"
	current := start
	steps := 0
	for ; current != goal; steps++ {
		currentNode := nodes[current]
		d := directions[steps%len(directions)]
		current = currentNode.neighbors[d]
	}
	return fmt.Sprintf("%d", steps)
}

func getMinStep(cycles []int) int {
	fmt.Println(cycles)
	minSteps := intmath.Lcm(cycles[0], cycles[1])
	for i := 2; i < len(cycles); i++ {
		minSteps = intmath.Lcm(minSteps, cycles[i])
	}
	return minSteps
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	directions, nodes := parseInput(input)
	starts := maps.Filter(nodes, func(_ string, n *node) bool { return n.isStart() })
	cycles := maps.MapToArray(starts, func(_ string, n *node) int { return n.getCycleLength(directions, nodes) })
	minSteps := getMinStep(cycles)
	return fmt.Sprintf("%d", minSteps)
}

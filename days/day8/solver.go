package day8

import (
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
	var directions []dir
	for _, c := range line {
		directions = append(directions, dir(c))
	}
	return directions
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
	current := nodes[n.label]
	counter := 0
	for ; !current.isGoal(); counter++ {
		currentNode := nodes[current.label]
		d := directions[counter%len(directions)]
		current = nodes[currentNode.neighbors[d]]
	}
	return counter
}

func parseNode(line string) *node {
	parts := strings.Split(line, " = ")
	label := parts[0]
	neighborsString := strings.Trim(parts[1], "()")
	neighborsList := strings.Split(neighborsString, ", ")
	neighbors := make(map[dir]string)
	dirs := []dir{left, right}
	for i, neighbor := range neighborsList {
		neighbors[dirs[i]] = neighbor
	}
	return &node{label, neighbors}
}

func parseInput(input string) ([]dir, map[string]*node) {
	parts := strings.Split(input, "\n\n")
	directions := parseDirections(parts[0])
	nodes := make(map[string]*node)
	lines := strings.Split(parts[1], "\n")
	for _, line := range lines {
		n := parseNode(line)
		nodes[n.label] = n
	}
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

func Filter(nodes map[string]*node, f func(*node) bool) []*node {
	filtered := make([]*node, 0)
	for _, n := range nodes {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func Map[T, U any](nodes []T, f func(T) U) []U {
	mapped := make([]U, 0)
	for _, n := range nodes {
		mapped = append(mapped, f(n))
	}
	return mapped
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func getMinStep(cycles []int) int {
	minSteps := lcm(cycles[0], cycles[1])
	for i := 2; i < len(cycles); i++ {
		minSteps = lcm(minSteps, cycles[i])
	}
	return minSteps
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	directions, nodes := parseInput(input)
	starts := Filter(nodes, func(n *node) bool { return n.isStart() })
	cycles := Map(starts, func(n *node) int { return n.getCycleLength(directions, nodes) })
	minSteps := getMinStep(cycles)
	return fmt.Sprintf("%d", minSteps)
}

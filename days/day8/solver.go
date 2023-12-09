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

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}

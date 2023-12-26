package day8

import (
	"aoc_2023/config"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day    = 8
	result any
)

func TestDay8Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "2"},
		{"Test 2", *config.NewTestInput(2), nil, "6"},
		{"Real", *config.NewRealInput(), nil, "16409"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart1(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func BenchmarkDebug(b *testing.B) {
	// input := config.NewRealInput()
	// cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	// solver := &Solver{}
	fileName := fmt.Sprintf("../../inputs/day%d/data%s.txt", 8, "")
	fileName, _ = filepath.Abs(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		directions, nodes := parseInput(lines)
		start := "AAA"
		goal := "ZZZ"
		current := start
		steps := 0
		for ; current != goal; steps++ {
			currentNode := nodes[current]
			d := directions[steps%len(directions)]
			current = currentNode.neighbors[d]
		}
		result = fmt.Sprintf("%d", steps)
	}
}

func BenchmarkDay8Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1(cfg.GetInputData())
	}
}

func TestDay8Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 3", *config.NewTestInput(3), nil, "6"},
		{"Real", *config.NewRealInput(), nil, "11795205644011"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart2(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func BenchmarkDay8Part2(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2(cfg.GetInputData())
	}
}

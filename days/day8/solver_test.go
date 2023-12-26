package day8

import (
	"aoc_2023/config"
	"aoc_2023/utils/maps"
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
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	// solver := &Solver{}
	directions, nodes := parseInput(cfg.GetInputData())
	starts := maps.Filter(nodes, func(_ string, n *node) bool { return n.isStart() })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = maps.MapToArray(starts, func(_ string, n *node) int { return n.getCycleLength(directions, nodes) })
	}
}

func BenchmarkDay8Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
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

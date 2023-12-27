package day2

import (
	"aoc_2023/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	day    = 2
	result any
)

func TestDay2Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), []any{"12", "13", "14"}, "8"},
		{"Real", *config.NewRealInput(), []any{"12", "13", "14"}, "2285"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart1(cfg.GetInputLines(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func BenchmarkDay2Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1(cfg.GetInputLines(), []any{"12", "13", "14"}...)
	}
}

func TestDay2Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "2286"},
		{"Real", *config.NewRealInput(), nil, "77021"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart2(cfg.GetInputLines(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func BenchmarkDay2Part2(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2(cfg.GetInputLines())
	}
}

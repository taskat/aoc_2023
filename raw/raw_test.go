package day_daynumber_

import (
	"aoc_2023/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	day    = 0
	result any
)

func Test_2023_Day_daynumber__Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, ""},
		{"Real", *config.NewRealInput(), nil, ""},
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

func Benchmark_2023_Day_daynumber__Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1(cfg.GetInputLines())
	}
}

func Test_2023_Day_daynumber__Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, ""},
		{"Real", *config.NewRealInput(), nil, ""},
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

func Benchmark_2023_Day_daynumber__Part2(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2(cfg.GetInputLines())
	}
}

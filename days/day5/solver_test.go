package day5

import (
	"aoc_2023/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	day    = 5
	result any
)

func Test_2023_Day4_Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "35"},
		{"Real", *config.NewRealInput(), nil, "389056265"},
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

func Benchmark_2023_Day5_Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1(cfg.GetInputLines())
	}
}

func Test_2023_Day4_Part2(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "46"},
		{"Test 2", *config.NewTestInput(2), nil, "15"},
		{"Real", *config.NewRealInput(), nil, "137516820"},
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

func Benchmark_2023_Day5_Part2(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2(cfg.GetInputLines())
	}
}

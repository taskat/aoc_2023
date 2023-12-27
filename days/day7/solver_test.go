package day7

import (
	"aoc_2023/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day    = 7
	result any
)

func TestDay7Part1(t *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "6440"},
		{"Real", *config.NewRealInput(), nil, "246912307"},
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

func BenchmarkDay7Part1(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart1(cfg.GetInputLines())
	}
}

func TestDay7Part2(b *testing.T) {
	testCases := []struct {
		name          string
		input         config.Input
		extraParams   []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "5905"},
		{"Real", *config.NewRealInput(), nil, "246894760"},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart2(cfg.GetInputLines(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func BenchmarkDay7Part2(b *testing.B) {
	input := config.NewRealInput()
	cfg := config.NewConfigForTest(config.NewConfig(day, 0, *input))
	solver := &Solver{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = solver.SolvePart2(cfg.GetInputLines())
	}
}

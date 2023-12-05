package day5

import (
	"aoc_2023/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var day = 5

func TestSolvePart1(t *testing.T) {
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
			solution := solver.SolvePart1(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func TestSolvePart2(t *testing.T) {
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
			solution := solver.SolvePart2(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

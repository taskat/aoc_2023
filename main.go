package main

import (
	"aoc_2023/config"
	"aoc_2023/days/day1"
	"aoc_2023/days/day2"
	"aoc_2023/days/day3"
	"aoc_2023/days/day4"
	"aoc_2023/days/day5"
	"aoc_2023/days/day6"
	"aoc_2023/days/day7"
	"aoc_2023/days/day8"
	"aoc_2023/solver"
	"fmt"
	"os"
)

type Config interface {
	GetDay() int
	GetPart() int
	GetInputType() string
	GetInputData() string
	GetExtraParams() []interface{}
}

func getSolver(cfg Config) solver.Solver {
	switch cfg.GetDay() {
	case 1:
		return &day1.Solver{}
	case 2:
		return &day2.Solver{}
	case 3:
		return &day3.Solver{}
	case 4:
		return &day4.Solver{}
	case 5:
		return &day5.Solver{}
	case 6:
		return &day6.Solver{}
	case 7:
		return &day7.Solver{}
	case 8:
		return &day8.Solver{}
	default:
		panic("Day not implemented yet")
	}
}

func solve(cfg Config, input string) string {
	solver := getSolver(cfg)
	if cfg.GetPart() == 1 {
		return solver.SolvePart1(input, cfg.GetExtraParams()...)
	} else {
		return solver.SolvePart2(input, cfg.GetExtraParams()...)
	}
}

func main() {
	cfg := config.ParseConfig()
	if cfg == nil {
		os.Exit(1)
	}
	input := cfg.GetInputData()
	fmt.Printf("Start solving day %d, part %d with %s input...\n", cfg.GetDay(), cfg.GetPart(), cfg.GetInputType())
	solution := solve(cfg, input)
	fmt.Printf("Solved! Solution is: %s\n", solution)
}

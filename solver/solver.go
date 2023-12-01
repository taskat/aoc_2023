package solver

type Solver interface {
	SolvePart1(input string, extraParams ...any) string
	SolvePart2(input string, extraParams ...any) string
}
package solver

type Solver interface {
	SolvePart1(lines []string, extraParams ...any) string
	SolvePart2(lines []string, extraParams ...any) string
}

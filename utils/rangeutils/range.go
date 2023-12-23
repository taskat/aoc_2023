package rangeutils

import "fmt"

type Range struct {
	from int
	to   int
}

func NewRange(from, to int) Range {
	if to < from {
		panic(fmt.Sprintf("Invalid range: %d-%d", from, to))
	}
	return Range{
		from: from,
		to:   to,
	}
}

func NewRangeWithLength(from, length int) Range {
	return Range{
		from: from,
		to:   from + length,
	}
}

func (r Range) Contains(value int) bool {
	return value >= r.from && value < r.to
}

func (r Range) ContainsRange(other Range) bool {
	return r.Contains(other.from) && r.Contains(other.to-1)
}

func (r Range) From() int {
	return r.from
}

func (r Range) HasIntersection(other Range) bool {
	return r.Contains(other.from) || r.Contains(other.to-1) || other.Contains(r.from) || other.Contains(r.to-1)
}

func (r Range) Length() int {
	return r.to - r.from
}

func (r Range) String() string {
	return fmt.Sprintf("From: %d, To: %d", r.from, r.to)
}

func (r Range) To() int {
	return r.to
}

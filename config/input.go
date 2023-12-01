package config

import "strconv"

type Input struct {
	real bool
	test int
}

func NewRealInput() *Input {
	return &Input{real: true}
}

func NewTestInput(test int) *Input {
	return &Input{test: test}
}

func (i *Input) String() string {
	if i.real {
		return ""
	} else {
		return strconv.Itoa(i.test)
	}
}
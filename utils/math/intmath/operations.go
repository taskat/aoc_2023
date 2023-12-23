package intmath

import "aoc_2023/utils/types"

func Abs[INT types.SignedInteger](number INT) INT {
	if number < 0 {
		return -number
	}
	return number
}

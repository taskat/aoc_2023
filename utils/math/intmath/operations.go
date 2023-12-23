package intmath

import "aoc_2023/utils/types"

func Abs[INT types.SignedInteger](number INT) INT {
	if number < 0 {
		return -number
	}
	return number
}

func Power[INT types.Integer](base, exponent INT) INT {
	var result INT = 1
	var i INT
	for i = 0; i < exponent; i++ {
		result *= base
	}
	return result
}

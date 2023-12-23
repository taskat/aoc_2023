package intmath

import "aoc_2023/utils/types"

func IsEven[INT types.Integer](number INT) bool {
	return number%2 == 0
}

func IsOdd[INT types.Integer](number INT) bool {
	return !IsEven(number)
}

func IsZero[INT types.Integer](number INT) bool {
	return number == 0
}

func IsNotZero[INT types.Integer](number INT) bool {
	return !IsZero(number)
}

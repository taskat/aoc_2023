package intmath

import (
	"aoc_2023/utils/types"
	"math"
)

func Abs[INT types.SignedInteger](number INT) INT {
	if number < 0 {
		return -number
	}
	return number
}

func Ceil[INT types.Integer](number float64) INT {
	return INT(math.Ceil(number))
}

func Floor[INT types.Integer](number float64) INT {
	return INT(math.Floor(number))
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Power[INT types.Integer](base, exponent INT) INT {
	var result INT = 1
	var i INT
	for i = 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func SolveQuadratic[INT types.Integer](a, b, c INT) (x1, x2 float64) {
	x1 = (-float64(b) + float64(math.Sqrt(float64(b*b-4*a*c)))) / (2 * float64(a))
	x2 = (-float64(b) - float64(math.Sqrt(float64(b*b-4*a*c)))) / (2 * float64(a))
	return
}

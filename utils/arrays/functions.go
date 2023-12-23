package arrays

import "aoc_2023/utils/types"

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Any[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if predicate(item) {
			return true
		}
	}
	return false
}

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range arr {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func ForEach[T any](arr []T, action func(T)) {
	for _, item := range arr {
		action(item)
	}
}

func Map[T any, U any](arr []T, mapper func(T) U) []U {
	var result []U
	for _, item := range arr {
		result = append(result, mapper(item))
	}
	return result
}

func MapToMap[T, VALUE any, KEY comparable](arr []T, mapper func(T) (KEY, VALUE)) map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	for _, item := range arr {
		key, value := mapper(item)
		result[key] = value
	}
	return result
}

func None[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if predicate(item) {
			return false
		}
	}
	return true
}

func Sum[T types.Summable](arr []T) T {
	var result T
	for _, item := range arr {
		result += item
	}
	return result
}

func Product[T types.Number](arr []T) T {
	result := T(1)
	for _, item := range arr {
		result *= item
	}
	return result
}

func Max[T types.Number](arr []T) T {
	if len(arr) == 0 {
		panic("Cannot find max of empty array")
	}
	result := arr[0]
	for _, item := range arr {
		if item > result {
			result = item
		}
	}
	return result
}

func Min[T types.Number](arr []T) T {
	if len(arr) == 0 {
		panic("Cannot find min of empty array")
	}
	result := arr[0]
	for _, item := range arr {
		if item < result {
			result = item
		}
	}
	return result
}
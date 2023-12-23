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

func Contains[T comparable](arr []T, item T) bool {
	for _, i := range arr {
		if i == item {
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

func Find[T any](arr []T, predicate func(T) bool) (T, bool) {
	for _, item := range arr {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func FindIndex[T any](arr []T, predicate func(T) bool) (int, bool) {
	for index, item := range arr {
		if predicate(item) {
			return index, true
		}
	}
	return -1, false
}

func ForEach[T any](arr []T, action func(T)) {
	for _, item := range arr {
		action(item)
	}
}

func Intersection[T comparable](arr1 []T, arr2 []T) []T {
	var result []T
	for _, item := range arr1 {
		if Contains(arr2, item) {
			result = append(result, item)
		}
	}
	return result
}

func Length[T any](arr []T) int {
	return len(arr)
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

func Remove[T any](arr []T, predicate func(T) bool) []T {
	return Filter(arr, func(item T) bool { return !predicate(item) })
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

func Max[T types.Real](arr []T) T {
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

func Min[T types.Real](arr []T) T {
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

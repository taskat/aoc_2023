package arrays

import "aoc_2023/utils/types"

func Accumulate[T, U any](arr []T, initial U, accumulator func(U, T) U) U {
	return Accumulate_i(arr, initial, func(init U, _ int, item T) U { return accumulator(initial, item) })
}

func Accumulate_i[T, U any](arr []T, initial U, accumulator func(U, int, T) U) U {
	result := initial
	for idx, item := range arr {
		result = accumulator(result, idx, item)
	}
	return result
}

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

func CountIf[T any](arr []T, predicate func(T) bool) int {
	counter := 0
	for _, item := range arr {
		if predicate(item) {
			counter++
		}
	}
	return counter
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
	_, result, ok := Find_i(arr, func(_ int, item T) bool { return predicate(item) })
	return result, ok
}

func Find_i[T any](arr []T, predicate func(int, T) bool) (int, T, bool) {
	for idx, item := range arr {
		if predicate(idx, item) {
			return idx, item, true
		}
	}
	var zero T
	return -1, zero, false
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

func Map[T, U any](arr []T, mapper func(T) U) []U {
	return Map_i(arr, func(_ int, item T) U { return mapper(item) })
}

func Map_i[T, U any](arr []T, mapper func(int, T) U) []U {
	var result []U
	for index, item := range arr {
		result = append(result, mapper(index, item))
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

func Pair[T, U any](arr1 []T, arr2 []U) []types.Pair[T, U] {
	if len(arr1) != len(arr2) {
		panic("Cannot pair arrays of different lengths")
	}
	result := make([]types.Pair[T, U], len(arr1))
	for i := 0; i < len(arr1); i++ {
		result[i] = *types.NewPair(arr1[i], arr2[i])
	}
	return result
}

func Remove[T any](arr []T, predicate func(T) bool) []T {
	return Filter(arr, func(item T) bool { return !predicate(item) })
}

func Split[T any](arr []T, predicate func(T) bool) [][]T {
	result := make([][]T, 0)
	current := make([]T, 0)
	for _, item := range arr {
		if predicate(item) {
			result = append(result, current)
			current = make([]T, 0)
		} else {
			current = append(current, item)
		}
	}
	result = append(result, current)
	return result
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
